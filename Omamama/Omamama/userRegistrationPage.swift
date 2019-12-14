//
//  ContentView.swift
//  Omamama
//
//  Created by shotaro-yamada on 2019/12/08.
//  Copyright © 2019 shotaro-yamada. All rights reserved.
//

import SwiftUI
import UIKit
import Foundation

struct userRegistrationView: View {
    @State var name = ""
    @State var firstPostalCode = ""
    @State var lastPostalCode = ""
    @State var address = ""
    @State var phoneNumber = ""
    
    var body: some View {
        VStack(){
            Text("ユーザー登録")
                .font(.largeTitle)
                .padding()
            Form{
                VStack(alignment: .leading, spacing: 20){
                    Text("氏名")
                        .font(.headline)
                        .padding()
                    TextField("氏名を入力してください", text: $name)
                        .textFieldStyle(RoundedBorderTextFieldStyle())
                    Text("郵便番号")
                        .font(.headline)
                        .padding()
                    HStack(){
                        TextField("000", text: $firstPostalCode)
                            .textFieldStyle(RoundedBorderTextFieldStyle())
                        Text("-")
                        TextField("0000", text: $lastPostalCode)
                            .textFieldStyle(RoundedBorderTextFieldStyle())
                        
                    }
                    .padding(.horizontal)
                    
                    Text("住所")
                        .font(.headline)
                        .padding()
                    TextField("", text: $address)
                        .textFieldStyle(RoundedBorderTextFieldStyle())
                }
                Text("電話番号")
                    .font(.headline)
                    .padding()
                TextField("", text: $phoneNumber)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                Spacer()
            }
            Button(action: {}){
                
                Text("次へ")
                    .font(.headline)
                    .frame(width: 100)
                    .accentColor(Color.white)
                    .padding()
                    .background(Color.green)
                    .cornerRadius(20)
            }
            .padding()
            .shadow(color: Color.gray, radius: 5.0, x: 5.0, y: 5.0)
        }
    }
}


struct userregistrationView_Previews: PreviewProvider {
    static var previews: some View {
        userRegistrationView()
    }
}
