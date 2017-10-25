package entity

// Meeting :

import (
 	"strings"
	// "Date"
	// "agenda-go-cli/loghelper"
)
type Meeting struct {
	Sponsor string
	Participators []string
	StartDate, EndDate Date
	Tittle string
}
func (m_meeting Meeting) init(t_Sponsor string, t_Participators []string, t_StartDate Date, t_EndDate Date,t_Tittle string) {
	m_meeting.Sponsor= t_Sponsor
	m_meeting.SetParticipator(t_Participators)
	m_meeting.StartDate.CopyDate(t_StartDate)
	m_meeting.EndDate.CopyDate(t_EndDate)
	m_meeting.Tittle= t_Tittle
}
func (m_meeting Meeting) CopyMeeting (t_meeting Meeting) {
	m_meeting.Sponsor= t_meeing.Sponsor
	m_meeting.SetParticipator(t_meeing.Participators)
	m_meeting.StartDate.CopyDate(t_meeing.StartDate)
	m_meeting.EndDate.CopyDate(t_meeing.EndDate)
	m_meeting.Tittle= t_meeing.Tittle
}
func (m_meeting Meeting) GetSponsor() string {
	return m_meeting.Sponsor
}

/**
* @brief set the sponsor of a meeting
* @param  the new sponsor string
*/
func (m_meeting Meeting) SetSponsor(t_sponsor string) {
	m_meeting.Sponsor = t_sponsor
}

/**
* @brief  get the participator of a meeting
* @return return a string indicate participator
*/
func (m_meeting Meeting) GetParticipator() []string {
    return m_meeting.Participators
}

/**
*   @brief set the new participator of a meeting
*   @param the new participator string
*/

func (m_meeting Meeting) SetParticipator(t_participators []string) {
	var length= t_participator.len()
	for i := 0; i < length; i++ {
		m_meeting.Participators[i]= t_participators[i]
	}
}

/**
* @brief get the startDate of a meeting
* @return return a string indicate startDate
*/
func (m_meeting Meeting) GetStartDate() Date {
	return m_meeting.StartDate
}

/**
* @brief  set the startDate of a meeting
* @param  the new startdate of a meeting
*/
func (m_meeting Meeting) SetStartDate(t_startTime Date) {
    m_meeting.StartDate.CopyDate(t_startTime)
}

/**
* @brief get the endDate of a meeting
* @return a date indicate the endDate
*/
func (m_meeting Meeting) GetEndDate() Date {
	return m_meeting.EndDate
}

/**
* @brief  set the endDate of a meeting
* @param  the new enddate of a meeting
*/
func (m_meeting Meeting) SetEndDate(t_endTime Date) {
	m_meeting.EndDate.CopyDate(t_endTime)
}

/**
* @brief get the title of a meeting
* @return a date title the endDate
*/
func (m_meeting Meeting) GetTitle() string {
	return m_meeting.Tittle
}

/**
* @brief  set the title of a meeting
* @param  the new title of a meeting
*/
func (m_meeting Meeting) SetTitle(t_title string) {
	m_meeting.Tittle = t_title
}

/**
* @brief check if the user take part in this meeting
* @param t_username the source username
* @return if the user take part in this meeting
*/
func (m_meeting Meeting) IsParticipator(t_username string) bool {
    var i int
	for i= 0; i< m_meeting.Participator.len(); i++ {
		if strings.EqualFold(m_meeting.Participators[i], t_username)== true {
	    	return true
		}
	}
	return false
}