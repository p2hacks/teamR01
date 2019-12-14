//
//  SendGiftViewController.swift
//  Omamama
//
//  Created by Naoki Noma on 2019/12/14.
//  Copyright © 2019 shotaro-yamada. All rights reserved.
//

import UIKit
import WebKit

final class SendGiftViewController: UIViewController,  WKNavigationDelegate{
//URL:http://images-jp.amazon.com/images/P/[ASIN,ISBN].[国コード].[画像の種類].jpg
    
    @IBOutlet final weak var BackgroundImage: UIImageView!
    
    @IBOutlet final weak var PresentImage: UIImageView!
    @IBOutlet final weak var PaidAnnounce: UIImageView!
    override func viewDidLoad() {
        super.viewDidLoad()
        let Asin:String = "test"
        let image:UIImage = GetImageByAmazon(url: "http://images-jp.amazon.com/images/P/" + Asin + ".09.MZZZZZZZ")
        MakeFilter(image: image)
        PresentImage.image = image
        PaidAnnounce.isHidden = true
        // Do any additional setup after loading the view.
        
        //noma
    }
    
    @IBAction func SendPresentButton(_ sender: Any) {
        //TODO: 画面遷移の実装
        //present
        DispatchQueue.main.asyncAfter(deadline: .now() + 0.5) {
            self.PaidAnnounce.isHidden = false
        }
        
    }
    
    func GetImageByAmazon(url: String)-> UIImage{
        let url = URL(string: url)
        do {
                let data = try Data(contentsOf: url!)
                return UIImage(data: data)!
            } catch let err {
                print("Error : \(err.localizedDescription)")
            }
            return UIImage()
    }
    
    func MakeFilter(image: UIImage){
        let BlurFilter = CIFilter(name: "CIGaussianBlur")
        BlurFilter!.setValue(image, forKey: kCIInputImageKey)
        
        let OutputImage: CIImage = BlurFilter!.outputImage!
        let OutputUIImage: UIImage = UIImage(ciImage: OutputImage)
        
        PresentImage.image = OutputUIImage
        PresentImage.setNeedsDisplay()
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
