(window.webpackJsonp=window.webpackJsonp||[]).push([[123],{684:function(d,l,t){"use strict";t.r(l);var c=t(1),Z=Object(c.a)({},(function(){var d=this,l=d.$createElement,t=d._self._c||l;return t("ContentSlotsDistributor",{attrs:{"slot-key":d.$parent.slotKey}},[t("h1",{attrs:{id:"integrating-ibc-middleware-into-a-chain"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#integrating-ibc-middleware-into-a-chain"}},[d._v("#")]),d._v(" Integrating IBC middleware into a chain")]),d._v(" "),t("p",[d._v("Learn how to integrate IBC middleware(s) with a base application to your chain. The following document only applies for Cosmos SDK chains.")]),d._v(" "),t("p",[d._v("If the middleware is maintaining its own state and/or processing SDK messages, then it should create and register its SDK module "),t("strong",[d._v("only once")]),d._v(" with the module manager in "),t("code",[d._v("app.go")]),d._v(".")]),d._v(" "),t("p",[d._v("All middleware must be connected to the IBC router and wrap over an underlying base IBC application. An IBC application may be wrapped by many layers of middleware, only the top layer middleware should be hooked to the IBC router, with all underlying middlewares and application getting wrapped by it.")]),d._v(" "),t("p",[d._v("The order of middleware "),t("strong",[d._v("matters")]),d._v(", function calls from IBC to the application travel from top-level middleware to the bottom middleware and then to the application. Function calls from the application to IBC goes through the bottom middleware in order to the top middleware and then to core IBC handlers. Thus the same set of middleware put in different orders may produce different effects.")]),d._v(" "),t("h2",{attrs:{id:"example-integration"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#example-integration"}},[d._v("#")]),d._v(" Example integration")]),d._v(" "),t("tm-code-block",{staticClass:"codeblock",attrs:{language:"go",base64:"Ly8gYXBwLmdvCgovLyBtaWRkbGV3YXJlIDEgYW5kIG1pZGRsZXdhcmUgMyBhcmUgc3RhdGVmdWwgbWlkZGxld2FyZSwgCi8vIHBlcmhhcHMgaW1wbGVtZW50aW5nIHNlcGFyYXRlIHNkay5Nc2cgYW5kIEhhbmRsZXJzCm13MUtlZXBlciA6PSBtdzEuTmV3S2VlcGVyKHN0b3JlS2V5MSkKbXczS2VlcGVyIDo9IG13My5OZXdLZWVwZXIoc3RvcmVLZXkzKQoKLy8gT25seSBjcmVhdGUgQXBwIE1vZHVsZSAqKm9uY2UqKiBhbmQgcmVnaXN0ZXIgaW4gYXBwIG1vZHVsZQovLyBpZiB0aGUgbW9kdWxlIG1haW50YWlucyBpbmRlcGVuZGVudCBzdGF0ZSBhbmQvb3IgcHJvY2Vzc2VzIHNkay5Nc2dzCmFwcC5tb2R1bGVNYW5hZ2VyID0gbW9kdWxlLk5ld01hbmFnZXIoCiAgLi4uCiAgbXcxLk5ld0FwcE1vZHVsZShtdzFLZWVwZXIpLAogIG13My5OZXdBcHBNb2R1bGUobXczS2VlcGVyKSwKICB0cmFuc2Zlci5OZXdBcHBNb2R1bGUodHJhbnNmZXJLZWVwZXIpLAogIGN1c3RvbS5OZXdBcHBNb2R1bGUoY3VzdG9tS2VlcGVyKQopCgptdzFJQkNNb2R1bGUgOj0gbXcxLk5ld0lCQ01vZHVsZShtdzFLZWVwZXIpCm13MklCQ01vZHVsZSA6PSBtdzIuTmV3SUJDTW9kdWxlKCkgLy8gbWlkZGxld2FyZTIgaXMgc3RhdGVsZXNzIG1pZGRsZXdhcmUKbXczSUJDTW9kdWxlIDo9IG13My5OZXdJQkNNb2R1bGUobXczS2VlcGVyKQoKc2NvcGVkS2VlcGVyVHJhbnNmZXIgOj0gY2FwYWJpbGl0eUtlZXBlci5OZXdTY29wZWRLZWVwZXIoJnF1b3Q7dHJhbnNmZXImcXVvdDspCnNjb3BlZEtlZXBlckN1c3RvbTEgOj0gY2FwYWJpbGl0eUtlZXBlci5OZXdTY29wZWRLZWVwZXIoJnF1b3Q7Y3VzdG9tMSZxdW90OykKc2NvcGVkS2VlcGVyQ3VzdG9tMiA6PSBjYXBhYmlsaXR5S2VlcGVyLk5ld1Njb3BlZEtlZXBlcigmcXVvdDtjdXN0b20yJnF1b3Q7KQoKLy8gTk9URTogSUJDIE1vZHVsZXMgbWF5IGJlIGluaXRpYWxpemVkIGFueSBudW1iZXIgb2YgdGltZXMgcHJvdmlkZWQgdGhleSB1c2UgYSBzZXBhcmF0ZQovLyBzY29wZWRLZWVwZXIgYW5kIHVuZGVybHlpbmcgcG9ydC4KCi8vIGluaXRpYWxpemUgYmFzZSBJQkMgYXBwbGljYXRpb25zCi8vIGlmIHlvdSB3YW50IHRvIGNyZWF0ZSB0d28gZGlmZmVyZW50IHN0YWNrcyB3aXRoIHRoZSBzYW1lIGJhc2UgYXBwbGljYXRpb24sCi8vIHRoZXkgbXVzdCBiZSBnaXZlbiBkaWZmZXJlbnQgc2NvcGVkS2VlcGVycyBhbmQgYXNzaWduZWQgZGlmZmVyZW50IHBvcnRzLgp0cmFuc2ZlcklCQ01vZHVsZSA6PSB0cmFuc2Zlci5OZXdJQkNNb2R1bGUodHJhbnNmZXJLZWVwZXIpCmN1c3RvbUlCQ01vZHVsZTEgOj0gY3VzdG9tLk5ld0lCQ01vZHVsZShjdXN0b21LZWVwZXIsICZxdW90O3BvcnRDdXN0b20xJnF1b3Q7KQpjdXN0b21JQkNNb2R1bGUyIDo9IGN1c3RvbS5OZXdJQkNNb2R1bGUoY3VzdG9tS2VlcGVyLCAmcXVvdDtwb3J0Q3VzdG9tMiZxdW90OykKCi8vIGNyZWF0ZSBJQkMgc3RhY2tzIGJ5IGNvbWJpbmluZyBtaWRkbGV3YXJlIHdpdGggYmFzZSBhcHBsaWNhdGlvbgovLyBOT1RFOiBzaW5jZSBtaWRkbGV3YXJlMiBpcyBzdGF0ZWxlc3MgaXQgZG9lcyBub3QgcmVxdWlyZSBhIEtlZXBlcgovLyBzdGFjayAxIGNvbnRhaW5zIG13MSAtJmd0OyBtdzMgLSZndDsgdHJhbnNmZXIKc3RhY2sxIDo9IG13MS5OZXdJQkNNaWRkbGV3YXJlKG13My5OZXdJQkNNaWRkbGV3YXJlKHRyYW5zZmVySUJDTW9kdWxlLCBtdzNLZWVwZXIpLCBtdzFLZWVwZXIpCi8vIHN0YWNrIDIgY29udGFpbnMgbXczIC0mZ3Q7IG13MiAtJmd0OyBjdXN0b20xCnN0YWNrMiA6PSBtdzMuTmV3SUJDTWlkZGxld2FyZShtdzIuTmV3SUJDTWlkZGxld2FyZShjdXN0b21JQkNNb2R1bGUxKSwgbXczS2VlcGVyKQovLyBzdGFjayAzIGNvbnRhaW5zIG13MiAtJmd0OyBtdzEgLSZndDsgY3VzdG9tMgpzdGFjazMgOj0gbXcyLk5ld0lCQ01pZGRsZXdhcmUobXcxLk5ld0lCQ01pZGRsZXdhcmUoY3VzdG9tSUJDTW9kdWxlMiwgbXcxS2VlcGVyKSkKCi8vIGFzc29jaWF0ZSBlYWNoIHN0YWNrIHdpdGggdGhlIG1vZHVsZU5hbWUgcHJvdmlkZWQgYnkgdGhlIHVuZGVybHlpbmcgc2NvcGVkS2VlcGVyCmliY1JvdXRlciA6PSBwb3J0dHlwZXMuTmV3Um91dGVyKCkKaWJjUm91dGVyLkFkZFJvdXRlKCZxdW90O3RyYW5zZmVyJnF1b3Q7LCBzdGFjazEpCmliY1JvdXRlci5BZGRSb3V0ZSgmcXVvdDtjdXN0b20xJnF1b3Q7LCBzdGFjazIpCmliY1JvdXRlci5BZGRSb3V0ZSgmcXVvdDtjdXN0b20yJnF1b3Q7LCBzdGFjazMpCmFwcC5JQkNLZWVwZXIuU2V0Um91dGVyKGliY1JvdXRlcikK"}})],1)}),[],!1,null,null,null);l.default=Z.exports}}]);