(self.webpackChunkgo_arcc=self.webpackChunkgo_arcc||[]).push([[674],{3905:function(e,t,r){"use strict";r.d(t,{Zo:function(){return u},kt:function(){return d}});var n=r(7294);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function a(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function l(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?a(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):a(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function c(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},a=Object.keys(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var i=n.createContext({}),p=function(e){var t=n.useContext(i),r=t;return e&&(r="function"==typeof e?e(t):l(l({},t),e)),r},u=function(e){var t=p(e.components);return n.createElement(i.Provider,{value:t},e.children)},s={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},m=n.forwardRef((function(e,t){var r=e.components,o=e.mdxType,a=e.originalType,i=e.parentName,u=c(e,["components","mdxType","originalType","parentName"]),m=p(r),d=o,k=m["".concat(i,".").concat(d)]||m[d]||s[d]||a;return r?n.createElement(k,l(l({ref:t},u),{},{components:r})):n.createElement(k,l({ref:t},u))}));function d(e,t){var r=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var a=r.length,l=new Array(a);l[0]=m;var c={};for(var i in t)hasOwnProperty.call(t,i)&&(c[i]=t[i]);c.originalType=e,c.mdxType="string"==typeof e?e:o,l[1]=c;for(var p=2;p<a;p++)l[p]=r[p];return n.createElement.apply(null,l)}return n.createElement.apply(null,r)}m.displayName="MDXCreateElement"},6979:function(e,t,r){"use strict";r.r(t),r.d(t,{frontMatter:function(){return c},contentTitle:function(){return i},metadata:function(){return p},toc:function(){return u},default:function(){return m}});var n=r(2122),o=r(9756),a=(r(7294),r(3905)),l=["components"],c={sidebar_position:1},i="GoArcc - Download",p={unversionedId:"Setup/Setup GoArcc",id:"Setup/Setup GoArcc",isDocsHomePage:!1,title:"GoArcc - Download",description:"To download GoArcc locally, you need to clone the repo.",source:"@site/docs/Setup/Setup GoArcc.md",sourceDirName:"Setup",slug:"/Setup/Setup GoArcc",permalink:"/GoArcc/docs/Setup/Setup GoArcc",editUrl:"https://github.com/facebook/docusaurus/edit/master/website/docs/Setup/Setup GoArcc.md",version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"Directory Definition",permalink:"/GoArcc/docs/Introduction/Directory Definition"},next:{title:"Service",permalink:"/GoArcc/docs/Setup/Service"}},u=[],s={toc:u};function m(e){var t=e.components,r=(0,o.Z)(e,l);return(0,a.kt)("wrapper",(0,n.Z)({},s,r,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"goarcc---download"},"GoArcc - Download"),(0,a.kt)("p",null,"To download GoArcc locally, you need to clone the repo."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"git clone https://github.com/deqode/GoArcc.git\n")),(0,a.kt)("p",null,"OR"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"go get github.com/deqode/GoArcc\n")),(0,a.kt)("h1",{id:"requirements"},"Requirements"),(0,a.kt)("p",null,"GoArcc depends on several tools and libraries-"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Go 1.16"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"https://docs.docker.com/install/"},"Docker")," 19.03+"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"https://docs.docker.com/compose/install/"},"Docker Compose")," 1.25+")),(0,a.kt)("p",null,"Below plugins are also required to generate code from a proto file. Check the usages.",(0,a.kt)("br",null)),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"https://github.com/grpc-ecosystem/grpc-gateway"},"gRPC-gateway")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"https://grpc.io/docs/protoc-installation/"},"protoc")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"https://github.com/envoyproxy/protoc-gen-validate"},"protoc-gen-validate")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"https://github.com/ysugimoto/grpc-graphql-gateway"},"grpc-graphql-gateway"))),(0,a.kt)("h1",{id:"running-the-goarcc"},"Running the GoArcc"),(0,a.kt)("p",null,"build the project using ",(0,a.kt)("inlineCode",{parentName:"p"},"make build"),(0,a.kt)("br",null)),(0,a.kt)("p",null,"run the binary ",(0,a.kt)("inlineCode",{parentName:"p"},"./bin/goarcc")),(0,a.kt)("p",null,"OR"),(0,a.kt)("p",null,"Run from docker-compose ",(0,a.kt)("inlineCode",{parentName:"p"},"docker-compose up")),(0,a.kt)("h1",{id:"hosts"},"Hosts"),(0,a.kt)("p",null,"GoArcc running servers hosts on different-2 ports."),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Jaeger UI: ",(0,a.kt)("inlineCode",{parentName:"li"},"http://localhost:16686")),(0,a.kt)("li",{parentName:"ul"},"Health Trace: ",(0,a.kt)("inlineCode",{parentName:"li"},"http://localhost:8083/health/")),(0,a.kt)("li",{parentName:"ul"},"Prometheus UI: ",(0,a.kt)("inlineCode",{parentName:"li"},"http://localhost:9090")),(0,a.kt)("li",{parentName:"ul"},"Prometheus UI Metrics: ",(0,a.kt)("inlineCode",{parentName:"li"},"http://localhost:9090/metrics")),(0,a.kt)("li",{parentName:"ul"},"Grpc Server: ",(0,a.kt)("inlineCode",{parentName:"li"},"http://localhost:8080")),(0,a.kt)("li",{parentName:"ul"},"Graphql Server: ",(0,a.kt)("inlineCode",{parentName:"li"},"http://localhost:8081")),(0,a.kt)("li",{parentName:"ul"},"Rest Server: ",(0,a.kt)("inlineCode",{parentName:"li"},"http://localhost:8082"))),(0,a.kt)("p",null,"Server configurations are in config.yml"))}m.isMDXComponent=!0}}]);