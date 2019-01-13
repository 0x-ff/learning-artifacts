
import Foundation

// Сборщик и парсер операнда - числа для выполнения операции
class OperandReader {
    
    // Была ли уже введена точка
    var isDotSet: Bool = false
    // Коллектор входящих символов
    var rawValue: String = "0"
    // Распарсенное число
    var value: Double = 0.0
    // Текст ошибки парсинга
    var lastError: String = ""
    
    init() {
        return
    }
    
    func setValue(_ val: Double) {
        value = val
        rawValue = String(value)
    }
    
    func getValue() -> Double {
        return value
    }
    
    func clearValue() {
        isDotSet = false
        rawValue = "0"
        value = 0.0
    }
    
    func appendDot() -> Bool {
        if isDotSet {
            return false
        }
        isDotSet = true
        rawValue.append(".")
        return true
    }
    
    func appendDigit(_ digit: Character) -> Bool {
        if digit == "0" && rawValue == "0" {
            return true
        }
        if rawValue == "0" {
            rawValue = ""
        }
        rawValue.append(digit)
        return true
    }
    
    func parse() -> Bool {
        let parsed = Double(rawValue)
        if parsed == nil {
            lastError = "error: \(rawValue)"
            return false
        }
        value = parsed!
        return true
    }
    
    func getRawValue() -> String {
        return rawValue
    }
    
    func getLastError() -> String {
        return lastError
    }
}
