
import Foundation

// Протокол числовых операций
protocol Operation {
    var name: String { get }
    var result: Double { get }
    var lastError: String { get }
    // Возвращает true если операция успешно выполнена и
    // false в случае ошибки
    func calculate() -> Bool
}

// Унарные операции sqrt, ln, log ...
protocol UnaryOperation: Operation {
    var operand: Double { get set }
}

// Бинарные операции + - * /
protocol BinaryOperation: Operation {
    var leftOperand: Double { get set }
    var rightOperand: Double { get set }
}

// Десятичный логарифм
class OperationLog: UnaryOperation {
    var name = "log"
    var operand: Double = 0.0
    var lastError: String = ""
    var result: Double = 0.0
    
    init() {
        return
    }
    
    func calculate() -> Bool {
        lastError = ""
        if operand > 0 {
            result = log10(operand)
            return true
        } else {
            lastError = "error: negative or zero value"
            return false
        }
    }
}

// Натуральный логарифм
class OperationLn: UnaryOperation {
    var name = "ln"
    var operand: Double = 0.0
    var lastError: String = ""
    var result: Double = 0.0
    
    init() {
        return
    }
    
    func calculate() -> Bool {
        lastError = ""
        if operand > 0 {
            result = log(operand)
            return true
        } else {
            lastError = "error: negative or zero value"
            return false
        }
    }
}

// Извлечение квадратного корня
class OperationSqrt: UnaryOperation {
    var name = "sqrt"
    var operand: Double = 0.0
    var lastError: String = ""
    var result: Double = 0.0
    
    init() {
        return
    }
    
    func calculate() -> Bool {
        lastError = ""
        if operand >= 0 {
            result = sqrt(operand)
            return true
        } else {
            lastError = "error: negative value"
            return false
        }
    }
}

// Сложение
class OperationPlus: BinaryOperation {
    var name = "plus"
    var leftOperand: Double = 0.0
    var rightOperand: Double = 0.0
    var lastError: String = ""
    var result: Double = 0.0
    
    init() {
        return
    }
    
    func calculate() -> Bool {
        result = leftOperand + rightOperand
        return true
    }
}

// Вычитание
class OperationMinus: BinaryOperation {
    var name = "minus"
    var leftOperand: Double = 0.0
    var rightOperand: Double = 0.0
    var lastError: String = ""
    var result: Double = 0.0
    
    init() {
        return
    }
    
    func calculate() -> Bool {
        result = leftOperand - rightOperand
        return true
    }
}

// Умножение
class OperationMul: BinaryOperation {
    var name = "mul"
    var leftOperand: Double = 0.0
    var rightOperand: Double = 0.0
    var lastError: String = ""
    var result: Double = 0.0
    
    init() {
        return
    }
    
    func calculate() -> Bool {
        result = leftOperand * rightOperand
        return true
    }
}

// Деление
class OperationDiv: BinaryOperation {
    var name = "div"
    var leftOperand: Double = 0.0
    var rightOperand: Double = 0.0
    var lastError: String = ""
    var result: Double = 0.0
    
    init() {
        return
    }
    
    func calculate() -> Bool {
        lastError = ""
        if rightOperand != 0 {
            result = leftOperand / rightOperand
            return true
        } else {
            lastError = "error: division by zero!"
            return false
        }
    }
}
