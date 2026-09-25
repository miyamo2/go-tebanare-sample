package usecase

import "github.com/miyamo2/go-tebanare-sample/pkg/collection"

// RecentlyViewedUseCase remembers the ids of the last few tasks a member
// looked at.
type RecentlyViewedUseCase struct {
	seen *collection.Stack[string]
}

// NewRecentlyViewedUseCase creates a RecentlyViewedUseCase.
func NewRecentlyViewedUseCase() *RecentlyViewedUseCase {
	return &RecentlyViewedUseCase{seen: collection.NewStack[string]()}
}

// Touch records that a member viewed the task with the given id.
func (uc *RecentlyViewedUseCase) Touch(taskID string) {
	uc.seen.Push(taskID)
}

// Count reports how many task views have been recorded. Not a getter: it
// calls a method on a field instead of returning a field directly.
func (uc *RecentlyViewedUseCase) Count() int { return uc.seen.Len() }

// Last returns the most recently viewed task id, or "" when none have been
// viewed yet. Not a getter: it does more than return a field.
func (uc *RecentlyViewedUseCase) Last() string {
	id, ok := uc.seen.Peek()
	if !ok {
		return ""
	}
	return id
}
