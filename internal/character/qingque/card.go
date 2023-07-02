package qingque

func (c *char) swap(pos1, pos2 int) {
	c.tiles[pos1], c.tiles[pos2] = c.tiles[pos2], c.tiles[pos1]
	c.suits[pos1], c.suits[pos2] = c.suits[pos2], c.suits[pos1]
}

// Functions if distinct suits aren't needed to save time and memory, for documentation read the comments on the functions including distinct suits
func (c *char) drawTile() {
	c.e2()
	if c.tiles[0] == 4 {
		return
	}

	s1, s2, s3 := c.tiles[0], c.tiles[1], c.tiles[2]
	startingTiles := s1 + s2 + s3
	drawn := c.engine.Rand().Intn(3)
	c.tiles[drawn] += 1
	switch {
	case c.tiles[1] > c.tiles[0]:
		c.swap(0, 1)
	case c.tiles[2] > c.tiles[0]:
		c.swap(0, 2)
	case c.tiles[2] > c.tiles[1]:
		c.swap(1, 2)
	}
	if startingTiles == 4 {
		c.discardTile()
	}
}

func (c *char) discardTile() {
	switch {
	case c.tiles[2] != 0:
		c.tiles[2] -= 1
	case c.tiles[1] != 0:
		c.tiles[1] -= 1
	default:
		c.tiles[0] -= 1
	}
}

/**
func (c *char) drawTile() string {
	c.e2()
	// Uses a suit-independent system to keep track of number of tiles.
	// This works because 4 of one suit functions the same way as 4 of another suit (difference lies in animations and UI)
	// and the probability of drawing each suit is the same (1/3), thus you can treat it as "symmetrical"
	// Ex: 2/1/1 or c.tiles = []int{2, 1, 1} means there's 2 of one suit and 1 of each of the other two suits.
	// c.suits represents the corresponding suit for each box, if c.suits = []string{"Wan", "Tong", "Tiao"} in the above case,
	// it means there is 2 Wan tiles, 1 Tong tile, and 1 Tiao tile. 3 basic suits are "Wan", "Tong", and "Tiao", the ult fish suit is "Yu",
	// no tile (as in when she only has <4 tiles) is "", and a blank tile (for skill purposes) I will just call "blank"

	// Returns blank tile if you're drawing with 4 of the same suit already
	if c.tiles[0] == 4 {
		return "blank"
	}
	// Gets number of tiles in each box and sums them for the total amount of tiles before anything is drawn
	s1, s2, s3 := c.tiles[0], c.tiles[1], c.tiles[2]
	startingTiles := s1 + s2 + s3
	// Random number generated to represent which box the newly drawn tile falls into
	drawn := c.engine.Rand().Intn(3)
	// If there is nothing in that box we need to pick a random suit that has not been used yet
	if c.tiles[drawn] == 0 {
		toUse := c.engine.Rand().Intn(len(c.unusedSuits))
		last := (len(c.unusedSuits) - 1)
		c.suits[drawn] = c.unusedSuits[toUse]
		c.unusedSuits[toUse] = c.unusedSuits[last]
		c.unusedSuits = c.unusedSuits[:last]
	}
	suit := c.suits[drawn]
	// Adds the tile then sorts the array in descending order (with some simplifications based on cases)
	// Only cases where stuff needs to be sorted are (0/0/1, 0/1/0, 1/0/1, 2/0/1, 2/3/0, 2/1/2). This sorts all of them properly
	c.tiles[drawn] += 1
	switch {
	case c.tiles[1] > c.tiles[0]:
		c.swap(0, 1)
	case c.tiles[2] > c.tiles[0]:
		c.swap(0, 2)
	case c.tiles[2] > c.tiles[1]:
		c.swap(1, 2)
	}
	if startingTiles == 4 {
		c.discardTile()
	}
	return suit
}
func (c *char) discardTile() string {
	// Pretty self explanatory, because our thing is in decreasing order just drop a tile from the last nonempty box
	// You should never be discarding with no tiles because you're only ever discarding at 5 tiles when gaining a tile from any source
	// or at 1-4 tiles due to QQ's basic (she always has at least 1 because she gains a tile at her own turn start)
	// For suits, when she has 2 or more of the same number of tiles from the suits and discards from one of them, she picks at random,
	// hence the random switching of suits into the position that is being discarded from.
	// Possible cases this happens are (1/1/0, 1/1/1, 2/2/0, 2/1/1, 3/1/1)
	// Then, removes a suit and repopulates unusedSuits if the amount of tiles of that suit reaches 0
	// Returns the suit that was discarded for the animations of skill and basic(iirc)

	suit := "Bug"
	switch {
	case c.tiles[2] != 0:
		if c.tiles[1] == c.tiles[2] {
			if c.tiles[2] == c.tiles[0] {
				switch c.engine.Rand().Intn(3) {
				case 0:
					c.suits[0] = c.suits[2]
				case 1:
					c.suits[1] = c.suits[2]
				}
			} else if c.engine.Rand().Intn(2) == 0 {
				c.suits[1] = c.suits[2]
			}
		}
		c.tiles[2] -= 1
		suit = c.suits[2]
		if c.tiles[2] == 0 {
			c.unusedSuits = c.unusedSuits[:(len(c.unusedSuits) + 1)]
			c.unusedSuits[len(c.unusedSuits)-1] = c.suits[2]
			c.suits[2] = ""
		}
	case c.tiles[1] != 0:
		if c.tiles[0] == c.tiles[1] && c.engine.Rand().Intn(2) == 0 {
			c.suits[0] = c.suits[1]
		}
		c.tiles[1] -= 1
		suit = c.suits[1]
		if c.tiles[1] == 0 {
			c.unusedSuits = c.unusedSuits[:(len(c.unusedSuits) + 1)]
			c.unusedSuits[len(c.unusedSuits)-1] = c.suits[1]
			c.suits[1] = ""
		}
	default:
		c.tiles[0] -= 1
		suit = c.suits[0]
		if c.tiles[0] == 0 {
			c.unusedSuits = c.unusedSuits[:(len(c.unusedSuits) + 1)]
			c.unusedSuits[len(c.unusedSuits)-1] = c.suits[0]
			c.suits[0] = ""
		}
	}
	return suit
}
**/
