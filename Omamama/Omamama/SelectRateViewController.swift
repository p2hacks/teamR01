//
//  SelectRateViewController.swift
//  Omamama
//
//  Created by Naoki Noma on 2019/12/15.
//  Copyright © 2019 shotaro-yamada. All rights reserved.
//

import UIKit

final class SelectRateViewController: UIViewController {

    override func viewDidLoad() {
        super.viewDidLoad()

        // Do any additional setup after loading the view.
    }
    @IBAction func LowRateButton(_ sender: Any) {
        let storyboard: UIStoryboard = self.storyboard!
        let selectgiftView = storyboard.instantiateViewController(withIdentifier: "selectgift") as! SelectGiftViewController
        self.present(selectgiftView, animated: true, completion: nil)
    }
    @IBAction func MiddleRateButton(_ sender: Any) {
    }
    @IBAction func HighRateButton(_ sender: Any) {
    }
    

    /*
    // MARK: - Navigation

    // In a storyboard-based application, you will often want to do a little preparation before navigation
    override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
        // Get the new view controller using segue.destination.
        // Pass the selected object to the new view controller.
    }
    */

}
