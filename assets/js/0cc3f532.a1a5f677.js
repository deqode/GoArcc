(self.webpackChunkgo_arcc=self.webpackChunkgo_arcc||[]).push([[388],{3905:function(e,t,r){"use strict";r.d(t,{Zo:function(){return u},kt:function(){return m}});var n=r(7294);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function a(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function c(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},i=Object.keys(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var s=n.createContext({}),p=function(e){var t=n.useContext(s),r=t;return e&&(r="function"==typeof e?e(t):a(a({},t),e)),r},u=function(e){var t=p(e.components);return n.createElement(s.Provider,{value:t},e.children)},l={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},d=n.forwardRef((function(e,t){var r=e.components,o=e.mdxType,i=e.originalType,s=e.parentName,u=c(e,["components","mdxType","originalType","parentName"]),d=p(r),m=o,f=d["".concat(s,".").concat(m)]||d[m]||l[m]||i;return r?n.createElement(f,a(a({ref:t},u),{},{components:r})):n.createElement(f,a({ref:t},u))}));function m(e,t){var r=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var i=r.length,a=new Array(i);a[0]=d;var c={};for(var s in t)hasOwnProperty.call(t,s)&&(c[s]=t[s]);c.originalType=e,c.mdxType="string"==typeof e?e:o,a[1]=c;for(var p=2;p<i;p++)a[p]=r[p];return n.createElement.apply(null,a)}return n.createElement.apply(null,r)}d.displayName="MDXCreateElement"},8215:function(e,t,r){"use strict";r.r(t),r.d(t,{frontMatter:function(){return c},contentTitle:function(){return s},metadata:function(){return p},toc:function(){return u},default:function(){return d}});var n=r(2122),o=r(9756),i=(r(7294),r(3905)),a=["components"],c={sidebar_position:1},s="What is GoArcc ?",p={unversionedId:"Introduction/About",id:"Introduction/About",isDocsHomePage:!1,title:"What is GoArcc ?",description:"When you start writing a Go project, GoArcc helps set up all the initial code boilerplate. Initial boilerplate code is how you organize your codebase and write multiple services. We have support for REST, Graphql and  gRPC.",source:"@site/docs/Introduction/About.md",sourceDirName:"Introduction",slug:"/Introduction/About",permalink:"/GoArcc/docs/Introduction/About",editUrl:"https://github.com/facebook/docusaurus/edit/master/website/docs/Introduction/About.md",version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"Introduction",permalink:"/GoArcc/docs/intro"},next:{title:"Prerequisite",permalink:"/GoArcc/docs/Introduction/Prerequisite"}},u=[{value:"Why we need GoArcc ?",id:"why-we-need-goarcc-",children:[]}],l={toc:u};function d(e){var t=e.components,r=(0,o.Z)(e,a);return(0,i.kt)("wrapper",(0,n.Z)({},l,r,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"what-is-goarcc-"},"What is GoArcc ?"),(0,i.kt)("p",null,"When you start writing a Go project, GoArcc helps set up all the initial code boilerplate. Initial boilerplate code is how you organize your codebase and write multiple services. We have support for ",(0,i.kt)("strong",{parentName:"p"},"REST"),", ",(0,i.kt)("strong",{parentName:"p"},"Graphql")," and  ",(0,i.kt)("strong",{parentName:"p"},"gRPC"),". "),(0,i.kt)("p",null,"We have implemented ",(0,i.kt)("strong",{parentName:"p"},"logging"),", ",(0,i.kt)("strong",{parentName:"p"},"tracing"),", ",(0,i.kt)("strong",{parentName:"p"},"health check"),", ",(0,i.kt)("strong",{parentName:"p"},"Jaeger"),", etc. to help developers write multiple services within a minute.\nIn short, ",(0,i.kt)("strong",{parentName:"p"},"GoArcc is a boilerplate setup codebase for any monolithic(Architecture) based web/mobile applications, which later converted into microservices(Architecture).")),(0,i.kt)("h2",{id:"why-we-need-goarcc-"},"Why we need GoArcc ?"),(0,i.kt)("p",null,"Nowadays, before setting up a new project, one needs to set up multiple things. Organizing the codebase is a primary concern for many and also takes a lot of time. As a result, our team tirelessly worked on it and eventually came up with a precise solution. Any web/mobile application whose backend is based upon monolith architecture can use our codebase and start writing services within seconds. We provide you with a setup of prewritten code that has multiple features like ",(0,i.kt)("strong",{parentName:"p"},"tracing"),", ",(0,i.kt)("strong",{parentName:"p"},"metrics"),", ",(0,i.kt)("strong",{parentName:"p"},"logging"),", ",(0,i.kt)("strong",{parentName:"p"},"health check"),", etc. This section will be covered later. "),(0,i.kt)("p",null,(0,i.kt)("strong",{parentName:"p"},"Note: At present, we only support Golang.")))}d.isMDXComponent=!0}}]);