//
//  SelectGiftViewController.swift
//  Omamama
//
//  Created by Naoki Noma on 2019/12/14.
//  Copyright © 2019 shotaro-yamada. All rights reserved.
//

import UIKit
import WebKit

final class SelectGiftViewController: UIViewController, WKNavigationDelegate {

    @IBOutlet private weak var GetAmazonPage: WKWebView!
    private var alertController : UIAlertController!
    
    //define black screen
    @IBOutlet weak var BlackOutView: UIView!{
        didSet{
            BlackOutView.backgroundColor = UIColor.black
            BlackOutView.alpha = 0.8
        }
    }
    @IBOutlet weak var BlackOutLabel_christmas: UILabel!{
        didSet{
            BlackOutLabel_christmas.textColor = UIColor.white
        }
    }
    @IBOutlet weak var BlackOutLabel_day: UILabel!{
        didSet{
            BlackOutLabel_day.textColor = UIColor.white
        }
    }
    @IBOutlet weak var BlackOutLabel_12_15: UILabel!{
        didSet{
            BlackOutLabel_12_15.textColor = UIColor.white
        }
    }
    
    @IBOutlet weak var BlackOutLabel_present: UILabel!{
        didSet{
            BlackOutLabel_present.textColor = UIColor.white
        }
    }
    
    @IBOutlet weak var BlackOutLabel_can: UILabel!{
        didSet{
            BlackOutLabel_can.textColor = UIColor.white
        }
    }
    
    
    
    override func viewDidLoad() {
        super.viewDidLoad()
        GetAmazonPage.navigationDelegate = self
        
        let myURL = URL(string: "https://amazon.co.jp")
        let myRequest = URLRequest(url: myURL!)
        GetAmazonPage.load(myRequest)
        BlackOutView.isHidden = true
    }
    
    @IBAction func SelectButton(_ sender: Any) {
        //アラート生成
        //部品のアラートを作る
        let alertController = UIAlertController(title: "確認", message: "このプレゼントで決定しますか？", preferredStyle: UIAlertController.Style.alert)
        //決定ボタン追加
        let okAction = UIAlertAction(title: "決定", style: UIAlertAction.Style.default, handler:{(action: UIAlertAction!) in
               DispatchQueue.main.asyncAfter(deadline: .now() + 0.5) {
                let activeUrl:URL? = self.GetAmazonPage.url
                let url = activeUrl?.absoluteString
                print(url!)
                //文字列のパース
                let arr:[String] = (url?.components(separatedBy: "/"))!
//TODO: send server
                print(arr[5])
                self.BlackOutView.isHidden = false
                }
               }
            )
        //戻るボタン追加
         let backAction = UIAlertAction(title: "戻る", style: UIAlertAction.Style.default, handler:nil)
        
        alertController.addAction(okAction)
        alertController.addAction(backAction)

        //アラートを表示する
         present(alertController, animated: true, completion: nil)
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
