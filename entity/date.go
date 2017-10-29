package entity

// Date :
import (
    "fmt"
    "strconv"
    "errors"
    // "agenda-go-cli/loghelper"
)
type Date struct {
    Year, Month, Day, Hour, Minute int
}

func (m_date Date) init(t_year, t_month, t_day, t_hour, t_minute int)  {
    m_date.Year= t_year
    m_date.Month= t_month
    m_date.Day= t_day
    m_date.Hour= t_hour
    m_date.Minute= t_minute
}
func (m_date Date) GetYear() int { 
    return m_date.Year
}
func (m_date Date) SetYear(t_year int) {
    m_date.Year= t_year
}
func (m_date Date) GetMonth() int { 
    return m_date.Month
}
func (m_date Date) SetMonth(t_month int) {
    m_date.Month= t_month
}
func (m_date Date) GetDay() int { 
    return m_date.Day
}
func (m_date Date) SetDay(t_day int) {
    m_date.Day= t_day
}
func (m_date Date) GetHour() int { 
    return m_date.Hour
}
func (m_date Date) SetHour(t_hour int) {
    m_date.Hour= t_hour
}
func (m_date Date) GetMinute() int { 
    return m_date.Minute
}
func (m_date Date) SetMinute(t_minute int) {
    m_date.Minute= t_minute
}

/**
*   @brief check whether the date is valid or not
*   @return the bool indicate valid or not
*/
func IsValid(t_date Date) bool {
    var current_year int= t_date.GetYear()
    var current_month int= t_date.GetMonth()
    var current_day int= t_date.GetDay()
    if current_year < 1000 || current_year > 9999 || current_month < 1 ||
        current_month > 12 || current_day < 1 || t_date.GetHour() < 0 ||
        t_date.GetHour() >= 24 || t_date.GetMinute() < 0 ||
        t_date.GetMinute() >= 60 {
        return false
    }
    if current_month == 1 || current_month == 3 || current_month == 5 ||
        current_month == 7 || current_month == 8 || current_month == 10 ||
        current_month == 12 {
        if current_day > 31 {
            return false
        }
     } else if current_month == 4 || current_month == 6 || current_month == 9 ||
               current_month == 11 {
        if current_day > 30 {
            return false
        }
     } else {
        //若年份为闰年，则2月29天
        if (current_year % 4 == 0 && current_year % 100 != 0) ||
            (current_year % 400 == 0) {
            if current_day > 29 {
                return false
            }
        } else {
            if current_day > 28 {
                return false
            }
        }
    }
    return true
}
/**
* @brief convert string to int
*/
func String2Int(s string) int {
    result,error :=strconv.Atoi(s)
    if error != nil{
        fmt.Println(error)
    }
    return result
}
/**
* @brief convert a string to date, if the format is not correct return
* 0000-00-00/00:00
* @return a date
*/
func StringToDate(t_dateString string) (Date, error) {
    var resultDate Date
    //检查字符串的格式是否正确．
    if (len(t_dateString) != 16) {
        return resultDate, errors.New("wrong")
    }
    var count int = 0 
    for count < len(t_dateString) {
        switch count {
            case 4:
                if t_dateString[4] != '-' {
                    return resultDate, errors.New("wrong")
                }
                break
            case 7:
                if t_dateString[7] != '-' {
                    return resultDate, errors.New("wrong")
                }
                break
            case 10:
                if t_dateString[10] != '/' {
                    return resultDate, errors.New("wrong")
                }
                break
            case 13:
                if t_dateString[13] != ':' {
                    return resultDate, errors.New("wrong")
                }
                break
            default:
                if t_dateString[count] < '0' || t_dateString[count] > '9' {
                    return resultDate, errors.New("wrong")
                }
        }
        count++
    }
    //若字符串格式没问题

    // resultDate.SetYear(String2Int(t_dateString[0:4]))
    resultDate.Year = String2Int(t_dateString[0:4])
    // resultDate.SetMonth(String2Int(t_dateString[5:7]))
    resultDate.Month = String2Int(t_dateString[5:7])
    // resultDate.SetDay(String2Int(t_dateString[8:10]))
    resultDate.Day = String2Int(t_dateString[8:10])
    // resultDate.SetHour(String2Int(t_dateString[11:13]))
    resultDate.Hour = String2Int(t_dateString[11:13])
    // resultDate.SetMinute(String2Int(t_dateString[14:]))
    resultDate.Minute = String2Int(t_dateString[14:])
    return resultDate,nil
}
/**
*   @brief convert the date to string, if result length is 1, add padding 0
*/
func Int2String(a int) string{
    var result_string string
    result_string=strconv.Itoa(a)  
    return result_string
}
/**
* @brief convert a date to string, if the format is not correct return
* 0000-00-00/00:00
*/

func DateToString(t_date Date) (string, error) {
    var dateString string = ""
    var wString string = ""
    var initTime string = "0000-00-00/00:00"
    //若date的格式错误，则返回初始时间串0000-00-00/00:00
    if !IsValid(t_date) {
        dateString = initTime
        return dateString,nil
    }
    dateString = Int2String(t_date.GetYear()) + "-" + Int2String(t_date.GetMonth()) +
        "-" + Int2String(t_date.GetDay()) + "/" + Int2String(t_date.GetHour()) +
        ":" + Int2String(t_date.GetMinute())
    if dateString != wString {
        return dateString, nil
    } else {
        return dateString, errors.New("wrong")
    }
}
/**
*  @brief overload the assign operator
*/
func (m_date Date) CopyDate (t_date Date) Date {
    m_date.SetYear(t_date.GetYear())
    m_date.SetMonth(t_date.GetMonth())
    m_date.SetDay(t_date.GetDay())
    m_date.SetHour(t_date.GetHour())
    m_date.SetMinute(t_date.GetMinute())
    return m_date
}

/**
* @brief check whether the CurrentDate is equal to the t_date
*/
func (m_date Date) IsSameDate(t_date Date) bool {
    return (t_date.GetYear() == m_date.GetYear() &&
            t_date.GetMonth() ==  m_date.GetMonth()&&
            t_date.GetDay() == m_date.GetDay()&&
            t_date.GetHour() ==  m_date.GetHour()&&
            t_date.GetMinute() == m_date.GetMinute())
}

/**
* @brief check whether the CurrentDate is  greater than the t_date
*/
func (m_date Date) MoreThan (t_date Date) bool {
    if m_date.Year > t_date.GetYear() {
        return true
    }
    if m_date.Year < t_date.GetYear() {
        return false
    }
    if m_date.Month > t_date.GetMonth() {
        return true
    }
    if m_date.Month < t_date.GetMonth() {
        return false
    }
    if m_date.Day > t_date.GetDay() {
        return true
    }
    if m_date.Day < t_date.GetDay() {
        return false
    }
    if m_date.Hour > t_date.GetHour() { 
        return true
    }
    if m_date.Hour < t_date.GetHour() {
        return false
    }
    if m_date.Minute > t_date.GetMinute() {
        return true
    }
    if m_date.Minute < t_date.GetMinute() {
        return false
    }
    return false
}
func (m_date Date) LessThan (t_date Date) bool {
    if m_date.IsSameDate(t_date)== false && m_date.MoreThan(t_date)== false {
        return true
    }
    return false
}
/**
* @brief check whether the CurrentDate is  greater or equal than the
* t_date
*/
func (m_date Date) GreateOrEqual(t_date Date) bool {
    return m_date.IsSameDate(t_date) || m_date.MoreThan(t_date)
}
/**
* @brief check whether the CurrentDate is  less than or equal to the
* t_date
*/
func (m_date Date) LessOrEqual(t_date Date) bool {
    return !m_date.MoreThan(t_date)
}