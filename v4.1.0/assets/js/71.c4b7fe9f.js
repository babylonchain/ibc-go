(window.webpackJsonp=window.webpackJsonp||[]).push([[71],{632:function(e,o,t){"use strict";t.r(o);var a=t(0),n=Object(a.a)({},(function(){var e=this,o=e.$createElement,t=e._self._c||o;return t("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[t("h1",{attrs:{id:"parameters"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#parameters"}},[e._v("#")]),e._v(" Parameters")]),e._v(" "),t("p",[e._v("The IBC transfer application module contains the following parameters:")]),e._v(" "),t("table",[t("thead",[t("tr",[t("th",[e._v("Key")]),e._v(" "),t("th",[e._v("Type")]),e._v(" "),t("th",[e._v("Default Value")])])]),e._v(" "),t("tbody",[t("tr",[t("td",[t("code",[e._v("SendEnabled")])]),e._v(" "),t("td",[e._v("bool")]),e._v(" "),t("td",[t("code",[e._v("true")])])]),e._v(" "),t("tr",[t("td",[t("code",[e._v("ReceiveEnabled")])]),e._v(" "),t("td",[e._v("bool")]),e._v(" "),t("td",[t("code",[e._v("true")])])])])]),e._v(" "),t("h2",{attrs:{id:"sendenabled"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#sendenabled"}},[e._v("#")]),e._v(" "),t("code",[e._v("SendEnabled")])]),e._v(" "),t("p",[e._v("The transfers enabled parameter controls send cross-chain transfer capabilities for all fungible tokens.")]),e._v(" "),t("p",[e._v("To prevent a single token from being transferred from the chain, set the "),t("code",[e._v("SendEnabled")]),e._v(" parameter to "),t("code",[e._v("true")]),e._v(" and then, depending on the Cosmos SDK version, do one of the following:")]),e._v(" "),t("ul",[t("li",[e._v("For Cosmos SDK v0.46.x or earlier, set the bank module's "),t("a",{attrs:{href:"https://github.com/cosmos/cosmos-sdk/blob/release/v0.46.x/x/bank/spec/05_params.md#sendenabled",target:"_blank",rel:"noopener noreferrer"}},[t("code",[e._v("SendEnabled")]),e._v(" parameter"),t("OutboundLink")],1),e._v(" for the denomination to "),t("code",[e._v("false")]),e._v(".")]),e._v(" "),t("li",[e._v("For Cosmos SDK versions above v0.46.x, set the bank module's "),t("code",[e._v("SendEnabled")]),e._v(" entry for the denomination to "),t("code",[e._v("false")]),e._v(" using "),t("code",[e._v("MsgSetSendEnabled")]),e._v(" as a governance proposal.")])]),e._v(" "),t("h2",{attrs:{id:"receiveenabled"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#receiveenabled"}},[e._v("#")]),e._v(" "),t("code",[e._v("ReceiveEnabled")])]),e._v(" "),t("p",[e._v("The transfers enabled parameter controls receive cross-chain transfer capabilities for all fungible tokens.")]),e._v(" "),t("p",[e._v("To prevent a single token from being transferred to the chain, set the "),t("code",[e._v("ReceiveEnabled")]),e._v(" parameter to "),t("code",[e._v("true")]),e._v(" and then, depending on the Cosmos SDK version, do one of the following:")]),e._v(" "),t("ul",[t("li",[e._v("For Cosmos SDK v0.46.x or earlier, set the bank module's "),t("a",{attrs:{href:"https://github.com/cosmos/cosmos-sdk/blob/release/v0.46.x/x/bank/spec/05_params.md#sendenabled",target:"_blank",rel:"noopener noreferrer"}},[t("code",[e._v("SendEnabled")]),e._v(" parameter"),t("OutboundLink")],1),e._v(" for the denomination to "),t("code",[e._v("false")]),e._v(".")]),e._v(" "),t("li",[e._v("For Cosmos SDK versions above v0.46.x, set the bank module's "),t("code",[e._v("SendEnabled")]),e._v(" entry for the denomination to "),t("code",[e._v("false")]),e._v(" using "),t("code",[e._v("MsgSetSendEnabled")]),e._v(" as a governance proposal.")])])])}),[],!1,null,null,null);o.default=n.exports}}]);