package entity

// User :


/**
* constructor with arguments
*/
type User struct {
	Name, Password, Email, Phone string
}
func (m_User User) init(t_Name, t_Password, t_Email, t_Phone string) {
	m_User.Name= t_Name
	m_User.Password= t_Password
	m_User.Email= t_Email
	m_User.Phone= t_Phone
}

/**
* @brief copy constructor
*/
func (m_User User) CopyUser(t_user User) {
	m_User.Name= t_user.Name
	m_User.Password= t_user.Password
	m_User.Email= t_user.Email
	m_User.Phone= t_user.Phone
}

/**
* @brief get the name of the user
* @return   return a string indicate the name of the user
*/
func (m_User User) GetName() string {
	return m_User.Name;
}

/**
* @brief set the name of the user
* @param   a string indicate the new name of the user
*/
func (m_User User) SetName(t_name string) {
	m_User.Name = t_name;
}

/**
* @brief get the password of the user
* @return   return a string indicate the password of the user
*/
func (m_User User) GetPassword() string {
	return m_User.Password;
}

/**
* @brief set the password of the user
* @param   a string indicate the new password of the user
*/
func (m_User User) SetPassword(t_password string) {
	m_User.Password = t_password;
}

/**
* @brief get the email of the user
* @return   return a string indicate the email of the user
*/
func (m_User User) GetEmail() string {
	return m_User.Email;
}

/**
* @brief set the email of the user
* @param   a string indicate the new email of the user
*/
func (m_User User) SetEmail(t_email string) {
	m_User.Email = t_email;
}

/**
* @brief get the phone of the user
* @return   return a string indicate the phone of the user
*/
func (m_User User) GetPhone() string {
	return m_User.Phone;
}

func (m_User User) SetPhone(t_phone string) {
	m_User.Phone = t_phone;
}
