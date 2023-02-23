(window.webpackJsonp=window.webpackJsonp||[]).push([[97],{660:function(e,t,o){"use strict";o.r(t);var n=o(1),s=Object(n.a)({},(function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[o("h1",{attrs:{id:"project-structure"}},[o("a",{staticClass:"header-anchor",attrs:{href:"#project-structure"}},[e._v("#")]),e._v(" Project structure")]),e._v(" "),o("p",[e._v("If you're not familiar with the overall module structure from the SDK modules, please check this "),o("a",{attrs:{href:"https://github.com/cosmos/cosmos-sdk/blob/main/docs/docs/building-modules/10-structure.md",target:"_blank",rel:"noopener noreferrer"}},[e._v("document"),o("OutboundLink")],1),e._v(" as prerequisite reading.")]),e._v(" "),o("p",[e._v("Every Interchain Standard (ICS) has been developed in its own package. The development team separated the IBC TAO (Transport, Authentication, Ordering) ICS specifications from the IBC application level specification. The following sections describe the architecture of the most relevant directories that comprise this repository.")]),e._v(" "),o("h2",{attrs:{id:"modules"}},[o("a",{staticClass:"header-anchor",attrs:{href:"#modules"}},[e._v("#")]),e._v(" "),o("code",[e._v("modules")])]),e._v(" "),o("p",[e._v("This folder contains implementations for the IBC TAO ("),o("code",[e._v("core")]),e._v("), IBC applications ("),o("code",[e._v("apps")]),e._v(") and light clients ("),o("code",[e._v("light-clients")]),e._v(").")]),e._v(" "),o("h3",{attrs:{id:"core"}},[o("a",{staticClass:"header-anchor",attrs:{href:"#core"}},[e._v("#")]),e._v(" "),o("code",[e._v("core")])]),e._v(" "),o("ul",[o("li",[o("code",[e._v("02-client")]),e._v(": This package is an implementation for Cosmos SDK-based chains of "),o("a",{attrs:{href:"https://github.com/cosmos/ibc/tree/main/spec/core/ics-002-client-semantics",target:"_blank",rel:"noopener noreferrer"}},[e._v("ICS 02"),o("OutboundLink")],1),e._v(". This implementation defines the types and methods needed to operate light clients tracking other chain's consensus state.")]),e._v(" "),o("li",[o("code",[e._v("03-connection")]),e._v(": This package is an implementation for Cosmos SDK-based chains of "),o("a",{attrs:{href:"https://github.com/cosmos/ibc/tree/main/spec/core/ics-003-connection-semantics",target:"_blank",rel:"noopener noreferrer"}},[e._v("ICS 03"),o("OutboundLink")],1),e._v(". This implementation defines the types and methods necessary to perform connection handshake between two chains.")]),e._v(" "),o("li",[o("code",[e._v("04-channel")]),e._v(": This package is an implementation for Cosmos SDK-based chains of "),o("a",{attrs:{href:"https://github.com/cosmos/ibc/tree/main/spec/core/ics-004-channel-and-packet-semantics",target:"_blank",rel:"noopener noreferrer"}},[e._v("ICS 04"),o("OutboundLink")],1),e._v(". This implementation defines the types and methods necessary to perform channel handshake between two chains and ensure correct packet sending flow.")]),e._v(" "),o("li",[o("code",[e._v("05-port")]),e._v(": This package is an implementation for Cosmos SDK-based chains of "),o("a",{attrs:{href:"https://github.com/cosmos/ibc/tree/main/spec/core/ics-005-port-allocation",target:"_blank",rel:"noopener noreferrer"}},[e._v("ICS 05"),o("OutboundLink")],1),e._v(". This implements the port allocation system by which modules can bind to uniquely named ports.")]),e._v(" "),o("li",[o("code",[e._v("23-commitment")]),e._v(": This package is an implementation for Cosmos SDK-based chains of "),o("a",{attrs:{href:"https://github.com/cosmos/ibc/tree/main/spec/core/ics-023-vector-commitments",target:"_blank",rel:"noopener noreferrer"}},[e._v("ICS 23"),o("OutboundLink")],1),e._v(". This implementation defines the functions required to prove inclusion or non-inclusion of particular values at particular paths in state.")]),e._v(" "),o("li",[o("code",[e._v("24-host")]),e._v(": This package is an implementation for Cosmos SDK-based chains of "),o("a",{attrs:{href:"https://github.com/cosmos/ibc/tree/main/spec/core/ics-024-host-requirements",target:"_blank",rel:"noopener noreferrer"}},[e._v("ICS 24"),o("OutboundLink")],1),e._v(".")])]),e._v(" "),o("h3",{attrs:{id:"apps"}},[o("a",{staticClass:"header-anchor",attrs:{href:"#apps"}},[e._v("#")]),e._v(" "),o("code",[e._v("apps")])]),e._v(" "),o("ul",[o("li",[o("code",[e._v("transfer")]),e._v(": This is the Cosmos SDK implementation of the "),o("a",{attrs:{href:"https://github.com/cosmos/ibc/tree/main/spec/app/ics-020-fungible-token-transfer",target:"_blank",rel:"noopener noreferrer"}},[e._v("ICS 20"),o("OutboundLink")],1),e._v(" protocol, which enables cross-chain fungible token transfers. For more information, read the "),o("RouterLink",{attrs:{to:"/apps/transfer/overview.html"}},[e._v("module's docs")])],1),e._v(" "),o("li",[o("code",[e._v("27-interchain-accounts")]),e._v(": This is the Cosmos SDK implementation of the "),o("a",{attrs:{href:"https://github.com/cosmos/ibc/tree/main/spec/app/ics-027-interchain-accounts",target:"_blank",rel:"noopener noreferrer"}},[e._v("ICS 27"),o("OutboundLink")],1),e._v(" protocol, which enables cross-chain account management built upon IBC. For more information, read the "),o("RouterLink",{attrs:{to:"/apps/interchain-accounts/overview.html"}},[e._v("module's documentation")]),e._v(".")],1),e._v(" "),o("li",[o("code",[e._v("29-fee")]),e._v(": This is the Cosmos SDK implementation of the "),o("a",{attrs:{href:"https://github.com/cosmos/ibc/tree/main/spec/app/ics-029-fee-payment",target:"_blank",rel:"noopener noreferrer"}},[e._v("ICS 29"),o("OutboundLink")],1),e._v(" middleware, which handles packet incentivisation and fee distribution on top of any ICS application protocol, enabling fee payment to relayer operators. For more information, read the "),o("RouterLink",{attrs:{to:"/middleware/ics29-fee/overview.html"}},[e._v("module's documentation")]),e._v(".")],1)]),e._v(" "),o("h3",{attrs:{id:"light-clients"}},[o("a",{staticClass:"header-anchor",attrs:{href:"#light-clients"}},[e._v("#")]),e._v(" "),o("code",[e._v("light-clients")])]),e._v(" "),o("ul",[o("li",[o("code",[e._v("06-solomachine")]),e._v(": This package implement the types for the Solo Machine light client specified in "),o("a",{attrs:{href:"https://github.com/cosmos/ibc/tree/main/spec/client/ics-006-solo-machine-client",target:"_blank",rel:"noopener noreferrer"}},[e._v("ICS 06"),o("OutboundLink")],1),e._v(".")]),e._v(" "),o("li",[o("code",[e._v("07-tendermint")]),e._v(": This package implement the types for the Tendermint consensus light client as specified in "),o("a",{attrs:{href:"https://github.com/cosmos/ibc/tree/main/spec/client/ics-007-tendermint-client",target:"_blank",rel:"noopener noreferrer"}},[e._v("ICS 07"),o("OutboundLink")],1),e._v(".")])]),e._v(" "),o("h2",{attrs:{id:"proto"}},[o("a",{staticClass:"header-anchor",attrs:{href:"#proto"}},[e._v("#")]),e._v(" "),o("code",[e._v("proto")])]),e._v(" "),o("p",[e._v("This folder contains all the Protobuf files used for")]),e._v(" "),o("ul",[o("li",[e._v("common message type definitions,")]),e._v(" "),o("li",[e._v("message type definitions related to genesis state,")]),e._v(" "),o("li",[o("code",[e._v("Query")]),e._v(" service and related message type definitions,")]),e._v(" "),o("li",[o("code",[e._v("Msg")]),e._v(" service and related message type definitions.")])]),e._v(" "),o("h2",{attrs:{id:"testing"}},[o("a",{staticClass:"header-anchor",attrs:{href:"#testing"}},[e._v("#")]),e._v(" "),o("code",[e._v("testing")])]),e._v(" "),o("p",[e._v("This package contains the implementation of the testing package used in unit and integration tests. Please read the "),o("RouterLink",{attrs:{to:"/testing/"}},[e._v("package's documentation")]),e._v(" for more information.")],1),e._v(" "),o("h2",{attrs:{id:"e2e"}},[o("a",{staticClass:"header-anchor",attrs:{href:"#e2e"}},[e._v("#")]),e._v(" "),o("code",[e._v("e2e")])]),e._v(" "),o("p",[e._v("This folder contains all the e2e tests of ibc-go. Please read the "),o("RouterLink",{attrs:{to:"/e2e/"}},[e._v("module's documentation")]),e._v(" for more information.")],1)])}),[],!1,null,null,null);t.default=s.exports}}]);