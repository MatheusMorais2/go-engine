package main

import (
	"bufio"
	"fmt"
	"os"
)

type StoryPage struct {
	text     string
	nextPage *StoryPage
	choice []*StoryPage
}

func main() {
	DeathPage := StoryPage{"Suddenly, you feel no more.", nil, nil}
	page1 := StoryPage{"The King and his men stole the Queen from her bed", nil, nil}
	page2 := StoryPage{"and bound her in her bones,", nil, nil}
	page3 := StoryPage{"The seas be ours, and by the Powers, where we will we roam!", nil, nil}
	firstChoice := []*StoryPage{
		{"You mourn, joining your brothers and sisters", nil, nil},
		&DeathPage,
	}
	page4 := StoryPage{"Do you cheer? y or n", nil, firstChoice}

	page1.nextPage = &page2
	page2.nextPage = &page3
	page3.nextPage = &page4


	firstChoice[0].addPage(&StoryPage{"You sob heavily upon hearing your fellow men do the same", nil, nil}, )
	firstChoice[1].addPage(&StoryPage{"You snark as the plan has come to fruition", nil, nil},)
	page1.removeNextPage()
	page1.addToTheEnd(&StoryPage{"And then the skeletons took his bones", nil, nil})



	page1.playStory()
}

func (page *StoryPage) playStory() error {

	if page == nil {
		end := fmt.Errorf("the end")
		fmt.Println(end)
		return end
	}

	if page.choice != nil {
		makeChoice(page.choice)
		return nil
	}

	fmt.Println(page.text)
	page.nextPage.playStory()

	return nil
}

func (storyPage *StoryPage) addPage(pageToAdd *StoryPage) {
	pageToAdd.nextPage = storyPage.nextPage
	storyPage.nextPage = pageToAdd
}

func (storyPage *StoryPage) addToTheEnd(pageToAdd *StoryPage) {
	for storyPage.nextPage != nil {
		storyPage = storyPage.nextPage
	}
	storyPage.nextPage = pageToAdd
}

func (storyPage *StoryPage) removeNextPage() {
	storyPage.nextPage = storyPage.nextPage.nextPage
}

func makeChoice(choicePage []*StoryPage)  error{
	readerOfInputs := bufio.NewReader(os.Stdin)
	fmt.Println("Do you cheer? y or n")
	input, _, _ := readerOfInputs.ReadRune()
	switch input {

	case 'n':
		choicePage[0].playStory()
		return nil
	case 'y':
		choicePage[1].playStory()
		return nil
	default:
		return nil
	}
}