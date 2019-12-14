//
//  SwiftUIView.swift
//  Omamama
//
//  Created by Itou Shingo on 2019/12/15.
//  Copyright © 2019 shotaro-yamada. All rights reserved.
//
import SwiftUI
import Foundation
import UIKit

struct cardInformationView: View {
    @State var cardNumber = ""
    @State var expirationMonth = ""
    @State var expirationYear = ""
    @State var securityCode = ""
    @State var cardName = ""
    
    var body: some View {
        VStack(spacing: 5){
            Text("クレジットカード登録")
                .font(.largeTitle)
            Spacer()
            VStack(alignment: .leading, spacing: 20){
                Spacer()
                Text("カード番号")
                    .font(.headline)
                TextField("000000000000", text: $cardNumber)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                    .frame(width: 150)
                Text("有効期限")
                    .font(.headline)
                    .padding()
                HStack(){
                    TextField("00", text: $expirationMonth)
                        .textFieldStyle(RoundedBorderTextFieldStyle())
                        .frame(width: 50)
                    Text("/")
                    TextField("00", text: $expirationYear)
                        .textFieldStyle(RoundedBorderTextFieldStyle())
                        .frame(width: 50)
                }
                Text("セキュリティーコード")
                    .font(.headline)
                TextField("000", text: $securityCode)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                    .frame(width: 80)
                Text("名義")
                    .font(.headline)
                TextField("カード名義人の氏名を入力", text: $cardName)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                    .frame(width: 250)
                Spacer()
            }
            Button(action: {}){
                Text("登録")
                    .font(.body)
                    .frame(width: 80)
                    .accentColor(Color.white)
                    .padding()
                    .background(Color.green)
                    .cornerRadius(20.0)
                
            }
            .shadow(color: Color.gray, radius: 5.0, x: 5.0, y: 5.0)
            .frame(width: 10, height: 70)
        }.padding()
    }
}

struct cardInformationView_Previews: PreviewProvider {
    static var previews: some View {
        cardInformationView()
    }
}
