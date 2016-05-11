package manage

import (
	"testing"

	"github.com/graffic/wanon/bot"
	"github.com/graffic/wanon/telegram"
	"github.com/graffic/wanon/test"
)

type mockQuoteMover struct {
	from string
	to   string
}

func (m *mockQuoteMover) Move(from string, to string) (int, error) {
	m.from = from
	m.to = to

	return 100, nil
}

type mockMessageSender struct {
	msg *telegram.SendMessage
}

func (m *mockMessageSender) SendMessage(msg *telegram.SendMessage) (*telegram.Message, error) {
	m.msg = msg
	return nil, nil
}

func TestHandle(t *testing.T) {
	mover := &mockQuoteMover{}
	handler := moveHandler{storage: mover}

	message := &telegram.Message{Chat: telegram.Chat{ID: 2}, MessageID: 3}
	sender := &mockMessageSender{}
	answer := telegram.NewAnswerBack(message, sender)
	context := &bot.MessageContext{Params: make(map[string]string), Message: answer}
	context.Params["from"] = "1"
	context.Params["to"] = "2"

	handler.Handle(context)

	if mover.from != "1" || mover.to != "2" {
		t.Error("Wrong storage call", mover)
	}

	if sender.msg.Text != "Moved 100 quotes" {
		t.Error("Wrong message", sender.msg.Text)
	}
}

func TestMoveStorage(t *testing.T) {
	dbHelper := test.NewGoejdbHelper(t, "TestMove")
	defer dbHelper.Cleanup()

	collOne := dbHelper.CreateColl("collOne")
	collTwo := dbHelper.CreateColl("collTwo")

	collOne.SaveJson("{}")
	collOne.SaveJson("{}")

	mover := NewEjdbQuoteMover(dbHelper.DB)
	mover.Move("collOne", "collTwo")

	amount, _ := collTwo.Count("{}")
	if amount != 2 {
		t.Error("The destination should have 2 items", amount)
	}

	collOne, _ = dbHelper.DB.GetColl("collOne")
	if collOne != nil {
		t.Error("collOne should not exist")
	}
}
