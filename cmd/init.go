/*
Copyright © 2020 Anoop B

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	upi "upi/init"
	keypair "upi/keys"
	"upi/qr"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise UPI payment generation",
	Long: `Initialises UPI payment generation by
taking VPA, name, among others as input fields`,
	Run: func(cmd *cobra.Command, args []string) {
		intent := upi.GenerateIntent()
		fmt.Println("unsigned intent:",intent.String())
		privateKey := keypair.GenerateRsaKeys(512)
		hashedIntent := upi.GetHash(intent.String())
		signature := upi.SignIntent(privateKey, hashedIntent)
		signedIntent:= upi.Concatenate(intent,signature)
		fmt.Println("Signed Intent:", signedIntent)
		qr.RenderString(signedIntent)
		response := upi.VerifySignature(&privateKey.PublicKey, hashedIntent, signature)
		fmt.Println("Verification passed:", response)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
