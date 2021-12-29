package dirac

type playerState struct {
	position, score int
}

func PlayDeterministic(p1Start, p2Start int) int {
	p1Position := p1Start - 1
	p2Position := p2Start - 1
	p1Score := 0
	p2Score := 0
	currentPlayer := 1
	die := DeterministicDie{}

	for p1Score < 1000 && p2Score < 1000 {
		roll := die.Roll() + die.Roll() + die.Roll()
		switch currentPlayer {
		case 1:
			p1Position = (p1Position + roll) % 10
			p1Score += p1Position + 1
			currentPlayer = 2
		case 2:
			p2Position = (p2Position + roll) % 10
			p2Score += p2Position + 1
			currentPlayer = 1
		}
	}

	if currentPlayer == 1 {
		return p1Score * die.rolls
	}

	return p2Score * die.rolls
}

func PlayQuantum(p1Start, p2Start int) int64 {
	p1Wins := int64(0)
	p2Wins := int64(0)
	diceDistribution := map[int]int64{
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	}
	p1States := map[playerState]int64{{p1Start - 1, 0}: 1}
	p1StateTotal := int64(1)
	p2States := map[playerState]int64{{p2Start - 1, 0}: 1}
	p2StateTotal := int64(1)

	for len(p1States) > 0 && len(p2States) > 0 {
		// Player 1
		newStates := make(map[playerState]int64)
		newStatesTotal := int64(0)
		for state, stateCount := range p1States {
			for roll, rollCount := range diceDistribution {
				newPosition := (state.position + roll) % 10
				newScore := state.score + newPosition + 1

				if newScore >= 21 {
					p1Wins += stateCount * rollCount * p2StateTotal
					continue
				}

				newStates[playerState{newPosition, newScore}] += stateCount * rollCount
				newStatesTotal += stateCount * rollCount
			}
		}
		p1States, newStates = newStates, make(map[playerState]int64)
		p1StateTotal, newStatesTotal = newStatesTotal, 0

		// Player 2
		if p1StateTotal == 0 {
			break
		}
		for state, stateCount := range p2States {
			for roll, rollCount := range diceDistribution {
				newPosition := (state.position + roll) % 10
				newScore := state.score + newPosition + 1

				if newScore >= 21 {
					p2Wins += stateCount * rollCount * p1StateTotal
					continue
				}

				newStates[playerState{newPosition, newScore}] += stateCount * rollCount
				newStatesTotal += stateCount * rollCount
			}
		}
		p2States = newStates
		p2StateTotal = newStatesTotal
	}

	if p1Wins > p2Wins {
		return p1Wins
	}

	return p2Wins
}
