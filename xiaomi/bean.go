package xiaomi

// SetPayload
func (m *Message) SetPayload(payLoad string) {
	m.Payload = payLoad
}

// SetRestrictedPackageName
func (m *Message) SetRestrictedPackageName(restrictedPackageName string) {
	m.RestrictedPackageName = restrictedPackageName
}

// SetTitle
func (m *Message) SetTitle(title string) {
	m.Title = title
}

// SetDescription
func (m *Message) SetDescription(description string) {
	m.Description = description
}

// SetNotifyType
func (m *Message) SetNotifyType(notifyType int) {
	m.NotifyType = notifyType
}

// SetNotifyId
func (m *Message) SetNotifyId(notifyId int) {
	m.NotifyId = notifyId
}

// SetRegistrationId
func (m *Message) SetRegistrationId(rid string) {
	m.RegistrationId = rid
}

// SetAlias
func (m *Message) SetAlias(alias string) {
	m.Alias = alias
}

// SetUserAccount
func (m *Message) SetUserAccount(account string) {
	m.UserAccount = account
}

// SetTopic
func (m *Message) SetTopic(topic string) {
	m.Topic = topic
}

// SetTopics
func (m *Message) SetTopics(topics string) {
	m.Topics = topics
}

// SetTopicOp
func (m *Message) SetTopicOp(topicOp string) {
	m.TopicOp = topicOp
}

// SetSoundUri
func (m *Message) SetSoundUri(soundUri string) {
	m.SoundUri = soundUri
}

// SetIntentUri
func (m *Message) SetIntentUri(intentUri string) {
	m.IntentUri = intentUri
}

// SetWebUri
func (m *Message) SetWebUri(webUri string) {
	m.WebUri = webUri
}
