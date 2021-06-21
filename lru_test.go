package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPutFirstNode(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	assert.Len(t, lruN.bucket, 1)
	assert.Equal(t, lruN.l.String(), "A=Apple,")
}

func TestPutKeyTwise(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	lruN.put("A", "Apple")
	assert.Len(t, lruN.bucket, 1)
	assert.Equal(t, lruN.l.String(), "A=Apple,")
}

func TestPutKeyTwiseDifferentVal(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	lruN.put("A", "ApplE")
	assert.Len(t, lruN.bucket, 1)
	assert.Equal(t, lruN.l.String(), "A=ApplE,")
}

func TestPutHalf(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	assert.Len(t, lruN.bucket, 2)
	assert.Equal(t, lruN.l.String(), "A=Apple,B=Ball,")
}

func TestPutHitFirstEleWithTwoElements(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	lruN.put("A", "ApplE")
	assert.Len(t, lruN.bucket, 2)
	assert.Equal(t, lruN.l.String(), "B=Ball,A=ApplE,")
}

func TestPutHitLastEleWithTwoElements(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	lruN.put("B", "Ball")
	assert.Len(t, lruN.bucket, 2)
	assert.Equal(t, lruN.l.String(), "A=Apple,B=Ball,")
}

func TestPutFull(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	lruN.put("C", "Cat")
	lruN.put("D", "Dog")
	assert.Len(t, lruN.bucket, 4)
	assert.Equal(t, lruN.l.String(), "A=Apple,B=Ball,C=Cat,D=Dog,")
}

func TestPutHitMiddleEle(t *testing.T) {
	lruN := NewLRU(5)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	lruN.put("C", "Cat")
	lruN.put("B", "BaLL")
	assert.Len(t, lruN.bucket, 3)
	assert.Equal(t, lruN.l.String(), "A=Apple,C=Cat,B=BaLL,")
}

func TestPutRemoveFirstRecentlyUsed(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	lruN.put("C", "Cat")
	lruN.put("D", "Dog")
	lruN.put("E", "Elephent")
	assert.Len(t, lruN.bucket, 4)
	assert.Equal(t, lruN.l.String(), "B=Ball,C=Cat,D=Dog,E=Elephent,")
}

func TestPutHitLastElement(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	lruN.put("C", "Cat")
	lruN.put("D", "Dog")
	lruN.put("D", "DOG")
	assert.Len(t, lruN.bucket, 4)
	assert.Equal(t, lruN.l.String(), "A=Apple,B=Ball,C=Cat,D=DOG,")
}

func TestPutRemoveSecondRecentlyUsed(t *testing.T) {
	lruN := NewLRU(4)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	lruN.put("A", "Apple")
	lruN.put("C", "Cat")
	lruN.put("D", "Dog")
	lruN.put("E", "Elephent")
	assert.Len(t, lruN.bucket, 4)
	assert.Equal(t, lruN.l.String(), "A=Apple,C=Cat,D=Dog,E=Elephent,")
}

func TestGetMissInEmpty(t *testing.T) {
	lruN := NewLRU(5)
	val := lruN.get("A")
	assert.Nil(t, val)
}

func TestGetMissInOneSize(t *testing.T) {
	lruN := NewLRU(5)
	lruN.put("A", "Apple")
	val := lruN.get("A")
	assert.Equal(t, *val, "Apple")
	assert.Equal(t, lruN.l.String(), "A=Apple,")
}

func TestGetHitFirstInPull(t *testing.T) {
	lruN := NewLRU(5)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	lruN.put("C", "Cat")
	lruN.put("D", "Dog")
	lruN.put("E", "Egg")
	val := lruN.get("A")
	assert.Equal(t, *val, "Apple")
	assert.Equal(t, lruN.l.String(), "B=Ball,C=Cat,D=Dog,E=Egg,A=Apple,")
}

func TestGetHitLastInPull(t *testing.T) {
	lruN := NewLRU(5)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	lruN.put("C", "Cat")
	lruN.put("D", "Dog")
	lruN.put("E", "Egg")
	val := lruN.get("E")
	assert.Equal(t, *val, "Egg")
	assert.Equal(t, lruN.l.String(), "A=Apple,B=Ball,C=Cat,D=Dog,E=Egg,")
}

func TestGetHitMiddelInPull(t *testing.T) {
	lruN := NewLRU(5)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	lruN.put("C", "Cat")
	lruN.put("D", "Dog")
	lruN.put("E", "Egg")
	val := lruN.get("C")
	assert.Equal(t, *val, "Cat")
	assert.Equal(t, lruN.l.String(), "A=Apple,B=Ball,D=Dog,E=Egg,C=Cat,")
}

func TestGetMissInPull(t *testing.T) {
	lruN := NewLRU(5)
	lruN.put("A", "Apple")
	lruN.put("B", "Ball")
	lruN.put("C", "Cat")
	lruN.put("D", "Dog")
	lruN.put("E", "Egg")
	val := lruN.get("F")
	assert.Nil(t, val)
	assert.Equal(t, lruN.l.String(), "A=Apple,B=Ball,C=Cat,D=Dog,E=Egg,")
}
