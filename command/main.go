package command

import "fmt"

type Command interface {
	Execute() error
}

type Gmail struct{}

func (g *Gmail) SendEmail(to string, subject string, body string) error {
	fmt.Println("Sending email to: ", to, " with subject: ", subject, " and body: ", body)
	return nil
}

type SendEmailParams struct {
	to      string
	subject string
	body    string
}

type SendEmailCommand struct {
	gmail  *Gmail
	params *SendEmailParams
}

func (c *SendEmailCommand) Execute() error {
	return c.gmail.SendEmail(c.params.to, c.params.subject, c.params.body)
}

type JobQueue struct {
	jobs []Command
}

func (q *JobQueue) Push(cmd Command) {
	q.jobs = append(q.jobs, cmd)
}

func (q *JobQueue) Run() {
	for _, job := range q.jobs {
		job.Execute()
	}
}

func Main() {
	gmail := &Gmail{}
	jobQueue := &JobQueue{}
	jobQueue.Push(&SendEmailCommand{
		gmail: gmail,
		params: &SendEmailParams{
			to:      "test@test.com",
			subject: "Test Subject",
			body:    "Test Body",
		}})
	jobQueue.Run()
}
