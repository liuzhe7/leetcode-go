/*
Author : Liu Zhe
Date   : 2023/1/14
*/

/**********************************************************************************

**********************************************************************************/

package stack

import "errors"

type StackRune struct {
	list []rune
}

func (s *StackRune) Insert(b rune) {
	s.list = append(s.list, b)
}

func (s *StackRune) Pop() (rune, error) {
	if len(s.list) == 0 {
		return '0', errors.New("stack is empty")
	}
	char := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return char, nil
}
