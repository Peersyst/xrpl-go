"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[716],{1169:(e,r,n)=>{n.r(r),n.d(r,{assets:()=>o,contentTitle:()=>d,default:()=>h,frontMatter:()=>t,metadata:()=>i,toc:()=>c});const i=JSON.parse('{"id":"keypairs","title":"keypairs","description":"Introduction","source":"@site/docs/keypairs.md","sourceDirName":".","slug":"/keypairs","permalink":"/docs/keypairs","draft":false,"unlisted":false,"editUrl":"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/keypairs.md","tags":[],"version":"current","sidebarPosition":5,"frontMatter":{"sidebar_position":5},"sidebar":"tutorialSidebar","previous":{"title":"address-codec","permalink":"/docs/address-codec"},"next":{"title":"xrpl"}}');var s=n(5105),a=n(3461);const t={sidebar_position:5},d="keypairs",o={},c=[{value:"Introduction",id:"introduction",level:2},{value:"Key components",id:"key-components",level:2},{value:"Supported algorithms",id:"supported-algorithms",level:2},{value:"crypto package",id:"crypto-package",level:3},{value:"API",id:"api",level:2},{value:"Key generation",id:"key-generation",level:3},{value:"GenerateSeed",id:"generateseed",level:4},{value:"DeriveKeypair",id:"derivekeypair",level:4},{value:"DeriveClassicAddress",id:"deriveclassicaddress",level:4},{value:"DeriveNodeAddress",id:"derivenodeaddress",level:4},{value:"Signing",id:"signing",level:3},{value:"Sign",id:"sign",level:4},{value:"Validate",id:"validate",level:4},{value:"Guides",id:"guides",level:2},{value:"How to generate a new random keypair",id:"how-to-generate-a-new-random-keypair",level:3},{value:"How to generate a new keypair from entropy",id:"how-to-generate-a-new-keypair-from-entropy",level:3}];function l(e){const r={a:"a",admonition:"admonition",code:"code",h1:"h1",h2:"h2",h3:"h3",h4:"h4",header:"header",li:"li",p:"p",pre:"pre",strong:"strong",ul:"ul",...(0,a.R)(),...e.components};return(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)(r.header,{children:(0,s.jsx)(r.h1,{id:"keypairs",children:"keypairs"})}),"\n",(0,s.jsx)(r.h2,{id:"introduction",children:"Introduction"}),"\n",(0,s.jsx)(r.p,{children:"The keypairs package provides a set of functions for generating and managing cryptographic keypairs. It includes functionality for creating new keypairs, deriving public keys from private keys, and verifying signatures."}),"\n",(0,s.jsxs)(r.p,{children:["This package is used internally by the ",(0,s.jsx)(r.code,{children:"xrpl"})," package to expose a ",(0,s.jsx)(r.a,{href:"/docs/xrpl/wallet",children:(0,s.jsx)(r.code,{children:"Wallet"})})," interface for easier wallet management. Nevertheless, it can be used independently of the ",(0,s.jsx)(r.code,{children:"xrpl"})," package for cryptographic operations."]}),"\n",(0,s.jsx)(r.h2,{id:"key-components",children:"Key components"}),"\n",(0,s.jsx)(r.p,{children:"This package works with the following key components from the XRP Ledger:"}),"\n",(0,s.jsxs)(r.ul,{children:["\n",(0,s.jsxs)(r.li,{children:[(0,s.jsx)(r.strong,{children:"Seed"}),": A base58-encoded string that represents a keypair."]}),"\n",(0,s.jsxs)(r.li,{children:[(0,s.jsx)(r.strong,{children:"Keypair"}),": A pair of a private and public key."]}),"\n",(0,s.jsxs)(r.li,{children:[(0,s.jsx)(r.strong,{children:"Address"}),": A base58-encoded string that represents an account."]}),"\n"]}),"\n",(0,s.jsxs)(r.p,{children:["To learn more about these components, you can check the ",(0,s.jsx)(r.a,{href:"https://xrpl.org/docs/concepts/accounts/cryptographic-keys",children:"official documentation"}),"."]}),"\n",(0,s.jsx)(r.h2,{id:"supported-algorithms",children:"Supported algorithms"}),"\n",(0,s.jsx)(r.p,{children:"Cryptographic algorithms supported by this package are:"}),"\n",(0,s.jsxs)(r.ul,{children:["\n",(0,s.jsx)(r.li,{children:"ed25519"}),"\n",(0,s.jsx)(r.li,{children:"secp256k1"}),"\n"]}),"\n",(0,s.jsxs)(r.p,{children:["Every function in the package that requires a cryptographic algorithm will accept any type that satisfies the ",(0,s.jsx)(r.code,{children:"KeypairCryptoAlg"})," interface. So, if desired, you can implement your own algorithm and use it in this package."]}),"\n",(0,s.jsxs)(r.p,{children:["However, the library already exports both algorithm getters that satisfy the ",(0,s.jsx)(r.code,{children:"KeypairCryptoAlg"})," and ",(0,s.jsx)(r.code,{children:"NodeDerivationCryptoAlg"})," interfaces. They're available under the package ",(0,s.jsx)(r.code,{children:"github.com/Peersyst/xrpl-go/pkg/crypto"}),", which exports both algorithm getters that satisfy the ",(0,s.jsx)(r.code,{children:"KeypairCryptoAlg"}),", ",(0,s.jsx)(r.code,{children:"NodeDerivationCryptoAlg"})," interfaces."]}),"\n",(0,s.jsx)(r.h3,{id:"crypto-package",children:"crypto package"}),"\n",(0,s.jsxs)(r.p,{children:["The ",(0,s.jsx)(r.code,{children:"crypto"})," package exports the following algorithm getters that satisfy the ",(0,s.jsx)(r.code,{children:"KeypairCryptoAlg"}),", ",(0,s.jsx)(r.code,{children:"NodeDerivationCryptoAlg"})," interfaces:"]}),"\n",(0,s.jsxs)(r.ul,{children:["\n",(0,s.jsx)(r.li,{children:(0,s.jsx)(r.code,{children:"ED25519()"})}),"\n",(0,s.jsx)(r.li,{children:(0,s.jsx)(r.code,{children:"SECP256K1()"})}),"\n"]}),"\n",(0,s.jsx)(r.p,{children:"You can use them to generate a seed or derive a keypair as the following example shows:"}),"\n",(0,s.jsx)(r.pre,{children:(0,s.jsx)(r.code,{className:"language-go",children:'seed, err := keypairs.GenerateSeed("", crypto.SECP256K1(), random.NewRandomizer())\n'})}),"\n",(0,s.jsx)(r.h2,{id:"api",children:"API"}),"\n",(0,s.jsx)(r.p,{children:"These are the functions available in this package:"}),"\n",(0,s.jsx)(r.pre,{children:(0,s.jsx)(r.code,{className:"language-go",children:"// Key generation\nfunc GenerateSeed(entropy string, alg interfaces.KeypairCryptoAlg, r interfaces.Randomizer) (string, error)\nfunc DeriveKeypair(seed string, validator bool) (private, public string, err error)\nfunc DeriveClassicAddress(pubKey string) (string, error)\nfunc DeriveNodeAddress(pubKey string, alg interfaces.NodeDerivationCryptoAlg) (string, error)\n\n// Signing\nfunc Sign(msg, privKey string) (string, error)\nfunc Validate(msg, pubKey, sig string) (bool, error)\n"})}),"\n",(0,s.jsx)(r.p,{children:"They can be split into two groups:"}),"\n",(0,s.jsxs)(r.ul,{children:["\n",(0,s.jsx)(r.li,{children:"Key generation: Functions that generate seeds and addresses."}),"\n",(0,s.jsx)(r.li,{children:"Signing: Functions that sign and validate messages."}),"\n"]}),"\n",(0,s.jsx)(r.h3,{id:"key-generation",children:"Key generation"}),"\n",(0,s.jsx)(r.h4,{id:"generateseed",children:"GenerateSeed"}),"\n",(0,s.jsx)(r.pre,{children:(0,s.jsx)(r.code,{className:"language-go",children:"func GenerateSeed(entropy string, alg interfaces.KeypairCryptoAlg, r interfaces.Randomizer) (string, error)\n"})}),"\n",(0,s.jsxs)(r.p,{children:["Generate a seed that can be used to generate keypairs. You can specify the entropy of the seed or let the function generate a random one (by passing an empty string as entropy and providing a randomizer) and use one of the supported algorithms to generate the seed. The result is a base58-encoded seed, which starts with the character ",(0,s.jsx)(r.code,{children:"s"}),"."]}),"\n",(0,s.jsx)(r.admonition,{type:"info",children:(0,s.jsxs)(r.p,{children:["A randomizer satisfies the ",(0,s.jsx)(r.code,{children:"Randomizer"})," interface. The ",(0,s.jsx)(r.code,{children:"random"})," package exports a ",(0,s.jsx)(r.code,{children:"NewRandomizer"})," function that returns a new randomizer."]})}),"\n",(0,s.jsx)(r.h4,{id:"derivekeypair",children:"DeriveKeypair"}),"\n",(0,s.jsx)(r.pre,{children:(0,s.jsx)(r.code,{className:"language-go",children:"func DeriveKeypair(seed string, validator bool) (private, public string, err error)\n"})}),"\n",(0,s.jsxs)(r.p,{children:["Derives a keypair (private and public keys) from a seed. If the ",(0,s.jsx)(r.code,{children:"validator"})," parameter is ",(0,s.jsx)(r.code,{children:"true"}),", the keypair will be a validator keypair; otherwise, it will be a user keypair. The result for both the private and public keys is a 33-byte hexadecimal string."]}),"\n",(0,s.jsx)(r.h4,{id:"deriveclassicaddress",children:"DeriveClassicAddress"}),"\n",(0,s.jsx)(r.pre,{children:(0,s.jsx)(r.code,{className:"language-go",children:"func DeriveClassicAddress(pubKey string) (string, error)\n"})}),"\n",(0,s.jsxs)(r.p,{children:["After deriving a keypair, you can derive the classic address from the public key. The result is a base58 encoded address, which starts with the character ",(0,s.jsx)(r.code,{children:"r"}),". If you're interested in X-Address derivation, the ",(0,s.jsx)(r.a,{href:"/docs/address-codec",children:(0,s.jsx)(r.code,{children:"address-codec"})})," package contains functions to encode and decode X-Addresses from and to classic addresses."]}),"\n",(0,s.jsx)(r.h4,{id:"derivenodeaddress",children:"DeriveNodeAddress"}),"\n",(0,s.jsx)(r.pre,{children:(0,s.jsx)(r.code,{className:"language-go",children:"func DeriveNodeAddress(pubKey string, alg interfaces.NodeDerivationCryptoAlg) (string, error)\n"})}),"\n",(0,s.jsxs)(r.p,{children:["Derives a node address from a public key. The result is a base58-encoded address, which starts with the character ",(0,s.jsx)(r.code,{children:"n"}),"."]}),"\n",(0,s.jsx)(r.h3,{id:"signing",children:"Signing"}),"\n",(0,s.jsx)(r.h4,{id:"sign",children:"Sign"}),"\n",(0,s.jsx)(r.pre,{children:(0,s.jsx)(r.code,{className:"language-go",children:"func Sign(msg, privKey string) (string, error)\n"})}),"\n",(0,s.jsxs)(r.p,{children:["Signs the provided message with the provided private key. To be able to sign a message, the private key must be a valid keypair and the message must be hex-encoded. The result is a hexadecimal string that represents the signature of the message. To verify the signature, you can use the ",(0,s.jsx)(r.code,{children:"Validate"})," function."]}),"\n",(0,s.jsx)(r.h4,{id:"validate",children:"Validate"}),"\n",(0,s.jsx)(r.pre,{children:(0,s.jsx)(r.code,{className:"language-go",children:"func Validate(msg, pubKey, sig string) (bool, error)\n"})}),"\n",(0,s.jsx)(r.p,{children:"Verifies a signature of a message. To be able to verify a signature, the public key must be valid, and the message and the signature must be hex-encoded. The result is a boolean value that indicates if the signature is valid or not."}),"\n",(0,s.jsx)(r.h2,{id:"guides",children:"Guides"}),"\n",(0,s.jsx)(r.h3,{id:"how-to-generate-a-new-random-keypair",children:"How to generate a new random keypair"}),"\n",(0,s.jsxs)(r.p,{children:["This example generates a new keypair using the ",(0,s.jsx)(r.code,{children:"SECP256K1"})," algorithm and a random entropy. It then derives a keypair from the seed and derives the classic address from the public key."]}),"\n",(0,s.jsx)(r.pre,{children:(0,s.jsx)(r.code,{className:"language-go",children:'package main\n\nimport (\n\t"fmt"\n\t"log"\n\n\t"github.com/Peersyst/xrpl-go/keypairs"\n\t"github.com/Peersyst/xrpl-go/pkg/crypto"\n\t"github.com/Peersyst/xrpl-go/pkg/random"\n)\n\nfunc main() {\n\tseed, err := keypairs.GenerateSeed("", crypto.SECP256K1(), random.NewRandomizer())\n\tif err != nil {\n\t\tlog.Fatal(err)\n\t}\n\n\tprivK, pubK, err := keypairs.DeriveKeypair(seed, false)\n\tif err != nil {\n\t\tlog.Fatal(err)\n\t}\n\n\taddr, err := keypairs.DeriveClassicAddress(pubK)\n\tif err != nil {\n\t\tlog.Fatal(err)\n\t}\n\n\tfmt.Println("Seed: ", seed)\n\tfmt.Println("Private Key: ", privK)\n\tfmt.Println("Public Key: ", pubK)\n\tfmt.Println("Address: ", addr)\n}\n'})}),"\n",(0,s.jsx)(r.h3,{id:"how-to-generate-a-new-keypair-from-entropy",children:"How to generate a new keypair from entropy"}),"\n",(0,s.jsxs)(r.p,{children:["This example generates a new keypair using the ",(0,s.jsx)(r.code,{children:"ED25519"})," algorithm and a provided entropy. Then, it derives the keypair and the address as the previous example."]}),"\n",(0,s.jsx)(r.pre,{children:(0,s.jsx)(r.code,{className:"language-go",children:'package main\n\nimport (\n\t"fmt"\n\t"log"\n\n\t"github.com/Peersyst/xrpl-go/keypairs"\n\t"github.com/Peersyst/xrpl-go/pkg/crypto"\n)\n\nfunc main() {\n\tseed, err := keypairs.GenerateSeed("ThisIsMyCustomEntropy", crypto.ED25519(), nil)\n\tif err != nil {\n\t\tlog.Fatal(err)\n\t}\n\n\tprivK, pubK, err := keypairs.DeriveKeypair(seed, false)\n\tif err != nil {\n\t\tlog.Fatal(err)\n\t}\n\n\taddr, err := keypairs.DeriveClassicAddress(pubK)\n\n\tfmt.Println("Seed: ", seed)\n\tfmt.Println("Private Key: ", privK)\n\tfmt.Println("Public Key: ", pubK)\n\tfmt.Println("Address: ", addr)\n}\n'})})]})}function h(e={}){const{wrapper:r}={...(0,a.R)(),...e.components};return r?(0,s.jsx)(r,{...e,children:(0,s.jsx)(l,{...e})}):l(e)}},3461:(e,r,n)=>{n.d(r,{R:()=>t,x:()=>d});var i=n(8101);const s={},a=i.createContext(s);function t(e){const r=i.useContext(a);return i.useMemo((function(){return"function"==typeof e?e(r):{...r,...e}}),[r,e])}function d(e){let r;return r=e.disableParentContext?"function"==typeof e.components?e.components(s):e.components||s:t(e.components),i.createElement(a.Provider,{value:r},e.children)}}}]);