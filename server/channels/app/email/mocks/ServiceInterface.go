// Code generated by mockery v2.10.4. DO NOT EDIT.

// Regenerate this file using `make email-mocks`.

package mocks

import (
	io "io"

	i18n "github.com/mattermost/mattermost-server/v6/server/platform/shared/i18n"

	mock "github.com/stretchr/testify/mock"

	model "github.com/mattermost/mattermost-server/v6/model"

	templates "github.com/mattermost/mattermost-server/v6/server/platform/shared/templates"

	throttled "github.com/throttled/throttled"
)

// ServiceInterface is an autogenerated mock type for the ServiceInterface type
type ServiceInterface struct {
	mock.Mock
}

// AddNotificationEmailToBatch provides a mock function with given fields: user, post, team
func (_m *ServiceInterface) AddNotificationEmailToBatch(user *model.User, post *model.Post, team *model.Team) *model.AppError {
	ret := _m.Called(user, post, team)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.User, *model.Post, *model.Team) *model.AppError); ok {
		r0 = rf(user, post, team)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// CreateVerifyEmailToken provides a mock function with given fields: userID, newEmail
func (_m *ServiceInterface) CreateVerifyEmailToken(userID string, newEmail string) (*model.Token, error) {
	ret := _m.Called(userID, newEmail)

	var r0 *model.Token
	if rf, ok := ret.Get(0).(func(string, string) *model.Token); ok {
		r0 = rf(userID, newEmail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userID, newEmail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMessageForNotification provides a mock function with given fields: post, translateFunc
func (_m *ServiceInterface) GetMessageForNotification(post *model.Post, translateFunc i18n.TranslateFunc) string {
	ret := _m.Called(post, translateFunc)

	var r0 string
	if rf, ok := ret.Get(0).(func(*model.Post, i18n.TranslateFunc) string); ok {
		r0 = rf(post, translateFunc)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetPerDayEmailRateLimiter provides a mock function with given fields:
func (_m *ServiceInterface) GetPerDayEmailRateLimiter() *throttled.GCRARateLimiter {
	ret := _m.Called()

	var r0 *throttled.GCRARateLimiter
	if rf, ok := ret.Get(0).(func() *throttled.GCRARateLimiter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*throttled.GCRARateLimiter)
		}
	}

	return r0
}

// InitEmailBatching provides a mock function with given fields:
func (_m *ServiceInterface) InitEmailBatching() {
	_m.Called()
}

// NewEmailTemplateData provides a mock function with given fields: locale
func (_m *ServiceInterface) NewEmailTemplateData(locale string) templates.Data {
	ret := _m.Called(locale)

	var r0 templates.Data
	if rf, ok := ret.Get(0).(func(string) templates.Data); ok {
		r0 = rf(locale)
	} else {
		r0 = ret.Get(0).(templates.Data)
	}

	return r0
}

// SendChangeUsernameEmail provides a mock function with given fields: newUsername, _a1, locale, siteURL
func (_m *ServiceInterface) SendChangeUsernameEmail(newUsername string, _a1 string, locale string, siteURL string) error {
	ret := _m.Called(newUsername, _a1, locale, siteURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(newUsername, _a1, locale, siteURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendCloudUpgradeConfirmationEmail provides a mock function with given fields: userEmail, name, trialEndDate, locale, siteURL, workspaceName, isYearly, embeddedFiles
func (_m *ServiceInterface) SendCloudUpgradeConfirmationEmail(userEmail string, name string, trialEndDate string, locale string, siteURL string, workspaceName string, isYearly bool, embeddedFiles map[string]io.Reader) error {
	ret := _m.Called(userEmail, name, trialEndDate, locale, siteURL, workspaceName, isYearly, embeddedFiles)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, string, string, bool, map[string]io.Reader) error); ok {
		r0 = rf(userEmail, name, trialEndDate, locale, siteURL, workspaceName, isYearly, embeddedFiles)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendCloudWelcomeEmail provides a mock function with given fields: userEmail, locale, teamInviteID, workSpaceName, dns, siteURL
func (_m *ServiceInterface) SendCloudWelcomeEmail(userEmail string, locale string, teamInviteID string, workSpaceName string, dns string, siteURL string) error {
	ret := _m.Called(userEmail, locale, teamInviteID, workSpaceName, dns, siteURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, string, string) error); ok {
		r0 = rf(userEmail, locale, teamInviteID, workSpaceName, dns, siteURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendDeactivateAccountEmail provides a mock function with given fields: _a0, locale, siteURL
func (_m *ServiceInterface) SendDeactivateAccountEmail(_a0 string, locale string, siteURL string) error {
	ret := _m.Called(_a0, locale, siteURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(_a0, locale, siteURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendDelinquencyEmail14 provides a mock function with given fields: _a0, locale, siteURL, planName
func (_m *ServiceInterface) SendDelinquencyEmail14(_a0 string, locale string, siteURL string, planName string) error {
	ret := _m.Called(_a0, locale, siteURL, planName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(_a0, locale, siteURL, planName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendDelinquencyEmail30 provides a mock function with given fields: _a0, locale, siteURL, planName
func (_m *ServiceInterface) SendDelinquencyEmail30(_a0 string, locale string, siteURL string, planName string) error {
	ret := _m.Called(_a0, locale, siteURL, planName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(_a0, locale, siteURL, planName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendDelinquencyEmail45 provides a mock function with given fields: _a0, locale, siteURL, planName, delinquencyDate
func (_m *ServiceInterface) SendDelinquencyEmail45(_a0 string, locale string, siteURL string, planName string, delinquencyDate string) error {
	ret := _m.Called(_a0, locale, siteURL, planName, delinquencyDate)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, string) error); ok {
		r0 = rf(_a0, locale, siteURL, planName, delinquencyDate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendDelinquencyEmail60 provides a mock function with given fields: _a0, locale, siteURL
func (_m *ServiceInterface) SendDelinquencyEmail60(_a0 string, locale string, siteURL string) error {
	ret := _m.Called(_a0, locale, siteURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(_a0, locale, siteURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendDelinquencyEmail7 provides a mock function with given fields: _a0, locale, siteURL, planName
func (_m *ServiceInterface) SendDelinquencyEmail7(_a0 string, locale string, siteURL string, planName string) error {
	ret := _m.Called(_a0, locale, siteURL, planName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(_a0, locale, siteURL, planName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendDelinquencyEmail75 provides a mock function with given fields: _a0, locale, siteURL, planName, delinquencyDate
func (_m *ServiceInterface) SendDelinquencyEmail75(_a0 string, locale string, siteURL string, planName string, delinquencyDate string) error {
	ret := _m.Called(_a0, locale, siteURL, planName, delinquencyDate)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, string) error); ok {
		r0 = rf(_a0, locale, siteURL, planName, delinquencyDate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendDelinquencyEmail90 provides a mock function with given fields: _a0, locale, siteURL
func (_m *ServiceInterface) SendDelinquencyEmail90(_a0 string, locale string, siteURL string) error {
	ret := _m.Called(_a0, locale, siteURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(_a0, locale, siteURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendEmailChangeEmail provides a mock function with given fields: oldEmail, newEmail, locale, siteURL
func (_m *ServiceInterface) SendEmailChangeEmail(oldEmail string, newEmail string, locale string, siteURL string) error {
	ret := _m.Called(oldEmail, newEmail, locale, siteURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(oldEmail, newEmail, locale, siteURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendEmailChangeVerifyEmail provides a mock function with given fields: newUserEmail, locale, siteURL, token
func (_m *ServiceInterface) SendEmailChangeVerifyEmail(newUserEmail string, locale string, siteURL string, token string) error {
	ret := _m.Called(newUserEmail, locale, siteURL, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(newUserEmail, locale, siteURL, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendGuestInviteEmails provides a mock function with given fields: team, channels, senderName, senderUserId, senderProfileImage, invites, siteURL, message, errorWhenNotSent, isSystemAdmin, isFirstAdmin
func (_m *ServiceInterface) SendGuestInviteEmails(team *model.Team, channels []*model.Channel, senderName string, senderUserId string, senderProfileImage []byte, invites []string, siteURL string, message string, errorWhenNotSent bool, isSystemAdmin bool, isFirstAdmin bool) error {
	ret := _m.Called(team, channels, senderName, senderUserId, senderProfileImage, invites, siteURL, message, errorWhenNotSent, isSystemAdmin, isFirstAdmin)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Team, []*model.Channel, string, string, []byte, []string, string, string, bool, bool, bool) error); ok {
		r0 = rf(team, channels, senderName, senderUserId, senderProfileImage, invites, siteURL, message, errorWhenNotSent, isSystemAdmin, isFirstAdmin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendInviteEmails provides a mock function with given fields: team, senderName, senderUserId, invites, siteURL, reminderData, errorWhenNotSent, isSystemAdmin, isFirstAdmin
func (_m *ServiceInterface) SendInviteEmails(team *model.Team, senderName string, senderUserId string, invites []string, siteURL string, reminderData *model.TeamInviteReminderData, errorWhenNotSent bool, isSystemAdmin bool, isFirstAdmin bool) error {
	ret := _m.Called(team, senderName, senderUserId, invites, siteURL, reminderData, errorWhenNotSent, isSystemAdmin, isFirstAdmin)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Team, string, string, []string, string, *model.TeamInviteReminderData, bool, bool, bool) error); ok {
		r0 = rf(team, senderName, senderUserId, invites, siteURL, reminderData, errorWhenNotSent, isSystemAdmin, isFirstAdmin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendInviteEmailsToTeamAndChannels provides a mock function with given fields: team, channels, senderName, senderUserId, senderProfileImage, invites, siteURL, reminderData, message, errorWhenNotSent, isSystemAdmin, isFirstAdmin
func (_m *ServiceInterface) SendInviteEmailsToTeamAndChannels(team *model.Team, channels []*model.Channel, senderName string, senderUserId string, senderProfileImage []byte, invites []string, siteURL string, reminderData *model.TeamInviteReminderData, message string, errorWhenNotSent bool, isSystemAdmin bool, isFirstAdmin bool) ([]*model.EmailInviteWithError, error) {
	ret := _m.Called(team, channels, senderName, senderUserId, senderProfileImage, invites, siteURL, reminderData, message, errorWhenNotSent, isSystemAdmin, isFirstAdmin)

	var r0 []*model.EmailInviteWithError
	if rf, ok := ret.Get(0).(func(*model.Team, []*model.Channel, string, string, []byte, []string, string, *model.TeamInviteReminderData, string, bool, bool, bool) []*model.EmailInviteWithError); ok {
		r0 = rf(team, channels, senderName, senderUserId, senderProfileImage, invites, siteURL, reminderData, message, errorWhenNotSent, isSystemAdmin, isFirstAdmin)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.EmailInviteWithError)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Team, []*model.Channel, string, string, []byte, []string, string, *model.TeamInviteReminderData, string, bool, bool, bool) error); ok {
		r1 = rf(team, channels, senderName, senderUserId, senderProfileImage, invites, siteURL, reminderData, message, errorWhenNotSent, isSystemAdmin, isFirstAdmin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendLicenseInactivityEmail provides a mock function with given fields: _a0, name, locale, siteURL
func (_m *ServiceInterface) SendLicenseInactivityEmail(_a0 string, name string, locale string, siteURL string) error {
	ret := _m.Called(_a0, name, locale, siteURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(_a0, name, locale, siteURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendLicenseUpForRenewalEmail provides a mock function with given fields: _a0, name, locale, siteURL, ctaTitle, ctaLink, ctaText, daysToExpiration
func (_m *ServiceInterface) SendLicenseUpForRenewalEmail(_a0 string, name string, locale string, siteURL string, ctaTitle string, ctaLink string, ctaText string, daysToExpiration int) error {
	ret := _m.Called(_a0, name, locale, siteURL, ctaTitle, ctaLink, ctaText, daysToExpiration)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, string, string, string, int) error); ok {
		r0 = rf(_a0, name, locale, siteURL, ctaTitle, ctaLink, ctaText, daysToExpiration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendMailWithEmbeddedFiles provides a mock function with given fields: to, subject, htmlBody, embeddedFiles, messageID, inReplyTo, references, category
func (_m *ServiceInterface) SendMailWithEmbeddedFiles(to string, subject string, htmlBody string, embeddedFiles map[string]io.Reader, messageID string, inReplyTo string, references string, category string) error {
	ret := _m.Called(to, subject, htmlBody, embeddedFiles, messageID, inReplyTo, references, category)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, map[string]io.Reader, string, string, string, string) error); ok {
		r0 = rf(to, subject, htmlBody, embeddedFiles, messageID, inReplyTo, references, category)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendMfaChangeEmail provides a mock function with given fields: _a0, activated, locale, siteURL
func (_m *ServiceInterface) SendMfaChangeEmail(_a0 string, activated bool, locale string, siteURL string) error {
	ret := _m.Called(_a0, activated, locale, siteURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool, string, string) error); ok {
		r0 = rf(_a0, activated, locale, siteURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendNoCardPaymentFailedEmail provides a mock function with given fields: _a0, locale, siteURL
func (_m *ServiceInterface) SendNoCardPaymentFailedEmail(_a0 string, locale string, siteURL string) error {
	ret := _m.Called(_a0, locale, siteURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(_a0, locale, siteURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendNotificationMail provides a mock function with given fields: to, subject, htmlBody
func (_m *ServiceInterface) SendNotificationMail(to string, subject string, htmlBody string) error {
	ret := _m.Called(to, subject, htmlBody)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(to, subject, htmlBody)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendPasswordChangeEmail provides a mock function with given fields: _a0, method, locale, siteURL
func (_m *ServiceInterface) SendPasswordChangeEmail(_a0 string, method string, locale string, siteURL string) error {
	ret := _m.Called(_a0, method, locale, siteURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(_a0, method, locale, siteURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendPasswordResetEmail provides a mock function with given fields: _a0, token, locale, siteURL
func (_m *ServiceInterface) SendPasswordResetEmail(_a0 string, token *model.Token, locale string, siteURL string) (bool, error) {
	ret := _m.Called(_a0, token, locale, siteURL)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, *model.Token, string, string) bool); ok {
		r0 = rf(_a0, token, locale, siteURL)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *model.Token, string, string) error); ok {
		r1 = rf(_a0, token, locale, siteURL)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendPaymentFailedEmail provides a mock function with given fields: _a0, locale, failedPayment, planName, siteURL
func (_m *ServiceInterface) SendPaymentFailedEmail(_a0 string, locale string, failedPayment *model.FailedPayment, planName string, siteURL string) (bool, error) {
	ret := _m.Called(_a0, locale, failedPayment, planName, siteURL)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, *model.FailedPayment, string, string) bool); ok {
		r0 = rf(_a0, locale, failedPayment, planName, siteURL)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, *model.FailedPayment, string, string) error); ok {
		r1 = rf(_a0, locale, failedPayment, planName, siteURL)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendRemoveExpiredLicenseEmail provides a mock function with given fields: ctaText, ctaLink, _a2, locale, siteURL
func (_m *ServiceInterface) SendRemoveExpiredLicenseEmail(ctaText string, ctaLink string, _a2 string, locale string, siteURL string) error {
	ret := _m.Called(ctaText, ctaLink, _a2, locale, siteURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, string) error); ok {
		r0 = rf(ctaText, ctaLink, _a2, locale, siteURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendSignInChangeEmail provides a mock function with given fields: _a0, method, locale, siteURL
func (_m *ServiceInterface) SendSignInChangeEmail(_a0 string, method string, locale string, siteURL string) error {
	ret := _m.Called(_a0, method, locale, siteURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(_a0, method, locale, siteURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendUserAccessTokenAddedEmail provides a mock function with given fields: _a0, locale, siteURL
func (_m *ServiceInterface) SendUserAccessTokenAddedEmail(_a0 string, locale string, siteURL string) error {
	ret := _m.Called(_a0, locale, siteURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(_a0, locale, siteURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendVerifyEmail provides a mock function with given fields: userEmail, locale, siteURL, token, redirect
func (_m *ServiceInterface) SendVerifyEmail(userEmail string, locale string, siteURL string, token string, redirect string) error {
	ret := _m.Called(userEmail, locale, siteURL, token, redirect)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, string) error); ok {
		r0 = rf(userEmail, locale, siteURL, token, redirect)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendWelcomeEmail provides a mock function with given fields: userID, _a1, verified, disableWelcomeEmail, locale, siteURL, redirect
func (_m *ServiceInterface) SendWelcomeEmail(userID string, _a1 string, verified bool, disableWelcomeEmail bool, locale string, siteURL string, redirect string) error {
	ret := _m.Called(userID, _a1, verified, disableWelcomeEmail, locale, siteURL, redirect)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, bool, bool, string, string, string) error); ok {
		r0 = rf(userID, _a1, verified, disableWelcomeEmail, locale, siteURL, redirect)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *ServiceInterface) Stop() {
	_m.Called()
}
