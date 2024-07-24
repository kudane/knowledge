package mail

type mailMessage struct {
	from string
	to   []string
	msg  []byte
}

func NewMailMessage() MailMessage {
	return &mailMessage{
		from: "",
		to:   []string{},
		msg:  nil,
	}
}

func (m mailMessage) From(f string) MailMessage {
	m.from = f
	return m
}

func (m mailMessage) To(f string) MailMessage {
	m.to = append(m.to, f)
	return m
}

func (m mailMessage) Body(b []byte) MailMessage {
	m.msg = b
	return m
}

func (m mailMessage) ToSend() MailBody {
	return MailBody{
		From: m.from,
		To:   m.to,
		msg:  m.msg,
	}
}
