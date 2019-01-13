//
//  ViewController.swift
//  calc
//
//  Created by Pavel Ivanisenko on 12/01/2019.
//  Copyright © 2019 Pavel Ivanisenko. All rights reserved.
//

import UIKit

class ViewController: UIViewController {

    @IBOutlet var output: [UITextField]!
    
    var calculator: Calculator = Calculator()
    
    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view, typically from a nib.
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }
    
    @IBAction func clickClearButton(_ sender: UIButton) {
        output[0].text = calculator
            .clearState()
            .getDisplayValue()
    }
    
    @IBAction func clickDotButton(_ sender: UIButton) {
        output[0].text = calculator
            .appendDot()
            .getDisplayValue()
    }
    
    @IBAction func click0Button(_ sender: UIButton) {
        output[0].text = calculator
            .appendDigit("0")
            .getDisplayValue()
    }
    
    @IBAction func click1Button(_ sender: UIButton) {
        output[0].text = calculator
            .appendDigit("1")
            .getDisplayValue()
    }
    
    @IBAction func click2Button(_ sender: UIButton) {
        output[0].text = calculator
            .appendDigit("2")
            .getDisplayValue()
    }
    
    @IBAction func click3Button(_ sender: UIButton) {
        output[0].text = calculator
            .appendDigit("3")
            .getDisplayValue()
    }
    
    @IBAction func click4Button(_ sender: UIButton) {
        output[0].text = calculator
            .appendDigit("4")
            .getDisplayValue()
    }
    
    @IBAction func click5Button(_ sender: UIButton) {
        output[0].text = calculator
            .appendDigit("5")
            .getDisplayValue()
    }
    
    @IBAction func click6Button(_ sender: UIButton) {
        output[0].text = calculator
            .appendDigit("6")
            .getDisplayValue()
    }
    
    @IBAction func click7Button(_ sender: UIButton) {
        output[0].text = calculator
            .appendDigit("7")
            .getDisplayValue()
    }
    
    @IBAction func click8Button(_ sender: UIButton) {
        output[0].text = calculator
            .appendDigit("8")
            .getDisplayValue()
    }
    
    @IBAction func click9button(_ sender: UIButton) {
        output[0].text = calculator
            .appendDigit("9")
            .getDisplayValue()
    }
    
    @IBAction func clickDivButton(_ sender: UIButton) {
        output[0].text = calculator
            .binaryOperation(calculator.OPERATION_DIV)
            .getDisplayValue()
    }
    
    @IBAction func clickMulButton(_ sender: UIButton) {
        output[0].text = calculator
            .binaryOperation(calculator.OPERATION_MUL)
            .getDisplayValue()
    }
    
    @IBAction func Minus(_ sender: UIButton) {
        output[0].text = calculator
            .binaryOperation(calculator.OPERATION_MINUS)
            .getDisplayValue()
    }
    
    @IBAction func clickPlusButton(_ sender: UIButton) {
        output[0].text = calculator
            .binaryOperation(calculator.OPERATION_PLUS)
            .getDisplayValue()
    }
    
    @IBAction func clickExecButton(_ sender: UIButton) {
        output[0].text = calculator
            .execute()
            .getDisplayValue()
    }
    
    @IBAction func clickSqrtButton(_ sender: UIButton) {
        output[0].text = calculator
            .calculateUnary(calculator.OPERATION_SQRT)
            .getDisplayValue()
    }
    
    @IBAction func clickLnButton(_ sender: UIButton) {
        output[0].text = calculator
            .calculateUnary(calculator.OPERATION_LN)
            .getDisplayValue()
    }
    
    @IBAction func clickLogButton(_ sender: UIButton) {
        output[0].text = calculator
            .calculateUnary(calculator.OPERATION_LOG)
            .getDisplayValue()
    }
}
