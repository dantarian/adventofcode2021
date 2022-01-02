package amphipod

import (
	"math"
	"sync"
	"sync/atomic"
)

func Solve(expanded bool) (int, error) {
	cells := createCells(expanded)
	initialState := createInitialState(cells)
	targetSuffix := "AABBCCDD"
	if expanded {
		targetSuffix += targetSuffix
	}

	// _, cost := initialState.solve(0)

	// return cost, nil

	initialMoves := initialState.legalMoves()

	results := make(chan int, len(initialMoves))

	var resultsCount int32
	var wg sync.WaitGroup

	for _, costedState := range initialMoves {
		wg.Add(1)
		newState := costedState.state
		cost := costedState.cost
		go func() {
			defer wg.Done()
			success, cost := newState.solve(cost)
			if success {
				atomic.AddInt32(&resultsCount, 1)
				results <- cost
			}
		}()
	}

	wg.Wait()

	minCost := math.MaxInt
	for i := 0; i < int(resultsCount); i++ {
		result := <-results
		if result < minCost {
			minCost = result
		}
	}
	return minCost, nil
}

func (s state) solve(costSoFar int) (bool, int) {
	minSuccessfulCost := math.MaxInt
	for _, costedState := range s.legalMoves() {
		if costedState.cost >= minSuccessfulCost {
			// No point doing further processing if we've already bust the budget.
			continue
		}

		if costedState.state.complete() {
			// We have a (potential) winner!
			minSuccessfulCost = costSoFar + costedState.cost
			continue
		}

		success, cost := costedState.state.solve(costSoFar + costedState.cost)
		if !success {
			continue
		}

		if cost < minSuccessfulCost {
			minSuccessfulCost = cost
		}
	}

	if minSuccessfulCost == math.MaxInt {
		return false, math.MaxInt
	}

	return true, minSuccessfulCost
}

type costedState struct {
	state state
	cost  int
}

func (s state) legalMoves() []costedState {
	neighbouringStates := []costedState{}
	for _, a := range s {
		if s.flavourComplete(a.flavour) {
			// All arthopods of this flavour are where they need to be.
			continue
		}

		if a.location.target && a.location.allowedFlavour == a.flavour {
			if len(a.location.linkedTo) == 1 {
				// We're at the back of our target room, so no moves to be made.
				continue
			}

			if s.amphipodDone(a) {
				// We're in the right place, so continue.
				continue
			}
		}

		baseCost := costPerCell(a.flavour)
		checkedCells := []*cell{a.location}
		cellsToCheck := make(map[*cell]int)
		for _, c := range a.location.linkedTo {
			cellsToCheck[c] = 1
		}
		for len(cellsToCheck) > 0 {
			for cellToCheck, distance := range cellsToCheck {
				checkedCells = append(checkedCells, cellToCheck)
				delete(cellsToCheck, cellToCheck)

				if cellToCheck.occupied(s) {
					// Can't enter or pass through an occupied space.
					continue
				}

				// Cell not occupied, so consider for onward travel.
				for _, nextCell := range cellToCheck.linkedTo {
					alreadyChecked := false
					for _, checkedCell := range checkedCells {
						if checkedCell == nextCell {
							alreadyChecked = true
							break
						}
					}

					if alreadyChecked {
						continue
					}

					cellsToCheck[nextCell] = distance + 1
				}

				if !cellToCheck.occupiable {
					continue
				}

				if cellToCheck.target {
					if cellToCheck.allowedFlavour != a.flavour {
						// Not our target.
						continue
					}

					if s.homeSpaceContainsForeigner(a.flavour) {
						// Can't enter.
						continue
					}

					if cellToCheck.id < a.location.id {
						// Don't want to move further out of our home room.
						continue
					}

					// Our home space is safe, so immediately move as far into it as possible.
					homeCell := cellToCheck
					additionalSteps := 0
					cellFound := false
					for !cellFound {
						if len(homeCell.linkedTo) == 1 {
							break
						}

						for _, nextCell := range homeCell.linkedTo {
							if nextCell.id < homeCell.id {
								continue
							}

							if nextCell.occupied(s) {
								cellFound = true
								break
							}

							homeCell = nextCell
							additionalSteps++
						}
					}
					return []costedState{{s.move(a.location, homeCell), (distance + additionalSteps) * baseCost}}
				}

				if !a.location.target {
					// Can't move between hallway spaces.
					continue
				}

				// All other options exhausted, this must be a viable state.
				neighbouringStates = append(neighbouringStates, costedState{s.move(a.location, cellToCheck), distance * baseCost})
			}
		}
	}

	return neighbouringStates
}

func (s state) amphipodDone(a amphipod) bool {
	if !a.location.target {
		return false
	}

	if a.location.allowedFlavour != a.flavour {
		return false
	}

	for _, cell := range a.location.linkedTo {
		if cell.id < a.location.id {
			continue
		}

		for _, a2 := range s {
			if a2.location.id == cell.id {
				return s.amphipodDone(a2)
			}
		}

		return false
	}

	return true
}

func (s state) move(oldCell, newCell *cell) state {
	newState := state{}
	for _, a := range s {
		if a.location == oldCell {
			newState = append(newState, amphipod{a.flavour, newCell})
			continue
		}

		newState = append(newState, a)
	}

	return newState
}

func (s state) flavourComplete(f flavour) bool {
	count := 0
	for _, a := range s {
		if a.flavour == f && a.location.target && a.location.allowedFlavour == f {
			count++
		}
	}

	return count == len(s)/4
}

func (s state) complete() bool {
	count := 0
	for _, a := range s {
		if a.location.target && a.location.allowedFlavour == a.flavour {
			count++
		}
	}

	return count == len(s)
}

func (s state) homeSpaceContainsForeigner(f flavour) bool {
	for _, a := range s {
		if a.flavour != f && a.location.allowedFlavour == f {
			return true
		}
	}
	return false
}

func (c *cell) occupied(s state) bool {
	for _, a := range s {
		if a.location == c {
			return true
		}
	}
	return false
}

func costPerCell(f flavour) int {
	switch f {
	case A:
		return 1
	case B:
		return 10
	case C:
		return 100
	case D:
		return 1000
	}
	return 0
}
