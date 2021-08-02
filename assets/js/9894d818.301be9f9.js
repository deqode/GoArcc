(self.webpackChunkgo_arcc=self.webpackChunkgo_arcc||[]).push([[715],{3905:function(e,t,n){"use strict";n.d(t,{Zo:function(){return p},kt:function(){return u}});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var s=r.createContext({}),c=function(e){var t=r.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},p=function(e){var t=c(e.components);return r.createElement(s.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},g=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,o=e.originalType,s=e.parentName,p=l(e,["components","mdxType","originalType","parentName"]),g=c(n),u=a,h=g["".concat(s,".").concat(u)]||g[u]||d[u]||o;return n?r.createElement(h,i(i({ref:t},p),{},{components:n})):r.createElement(h,i({ref:t},p))}));function u(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var o=n.length,i=new Array(o);i[0]=g;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l.mdxType="string"==typeof e?e:a,i[1]=l;for(var c=2;c<o;c++)i[c]=n[c];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}g.displayName="MDXCreateElement"},2262:function(e,t,n){"use strict";n.r(t),n.d(t,{frontMatter:function(){return l},contentTitle:function(){return s},metadata:function(){return c},toc:function(){return p},default:function(){return g}});var r=n(2122),a=n(9756),o=(n(7294),n(3905)),i=["components"],l={sidebar_position:3},s=void 0,c={unversionedId:"Introduction/Directory Definition",id:"Introduction/Directory Definition",isDocsHomePage:!1,title:"Directory Definition",description:"Before you start writing your services, you need to understand how GoArcc works and the definition of its directory.",source:"@site/docs/Introduction/Directory Definition.md",sourceDirName:"Introduction",slug:"/Introduction/Directory Definition",permalink:"/GoArcc/docs/Introduction/Directory Definition",editUrl:"https://github.com/facebook/docusaurus/edit/master/website/docs/Introduction/Directory Definition.md",version:"current",sidebarPosition:3,frontMatter:{sidebar_position:3},sidebar:"tutorialSidebar",previous:{title:"Prerequisite",permalink:"/GoArcc/docs/Introduction/Prerequisite"},next:{title:"GoArcc - Download",permalink:"/GoArcc/docs/Setup/Setup GoArcc"}},p=[{value:"<u>Config</u>",id:"config",children:[{value:"<u>How can you add a custom config?</u>",id:"how-can-you-add-a-custom-config",children:[]},{value:"<u>How you can pass enviroment variables?</u>",id:"how-you-can-pass-enviroment-variables",children:[]}]},{value:"<u>Db</u>",id:"db",children:[]},{value:"<u>Logger</u>",id:"logger",children:[{value:"<u>Why have we used zap logger? </u>",id:"why-have-we-used-zap-logger",children:[]},{value:"<u>Zap log level reference </u>",id:"zap-log-level-reference",children:[]}]},{value:"<u>Cmd</u>",id:"cmd",children:[{value:"<u>main.go</u>",id:"maingo",children:[]},{value:"<u>invokers.go</u>",id:"invokersgo",children:[]},{value:"<u>providers.go</u>",id:"providersgo",children:[]},{value:"<u>app.go</u>",id:"appgo",children:[]}]},{value:"<u> Servers </u>",id:"servers",children:[{value:"GRPC",id:"grpc",children:[]},{value:"REST",id:"rest",children:[]},{value:"GraphQL",id:"graphql",children:[]},{value:"HealthCheck",id:"healthcheck",children:[]},{value:"Jaeger",id:"jaeger",children:[]},{value:"Prometheus",id:"prometheus",children:[]}]}],d={toc:p};function g(e){var t=e.components,l=(0,a.Z)(e,i);return(0,o.kt)("wrapper",(0,r.Z)({},d,l,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("p",null,"Before you start writing your services, you need to understand how GoArcc works and the definition of its directory."),(0,o.kt)("p",null,"Let's dive deep into the definition and codebase."),(0,o.kt)("h1",{id:"directory-structure"},(0,o.kt)("u",null,"Directory Structure")),(0,o.kt)("p",null,"We follow a standard Go directory structure for our codebase so that the developers can easily adapt it. There are multiple directories and multiple libraries inside GoArcc. We will go through each one by one.\nThe diagram below shows multiple directories such as ",(0,o.kt)("strong",{parentName:"p"},"cmd"),", ",(0,o.kt)("strong",{parentName:"p"},"client"),", ",(0,o.kt)("strong",{parentName:"p"},"logger"),", ",(0,o.kt)("strong",{parentName:"p"},"protoc"),", etc."),(0,o.kt)("p",null,(0,o.kt)("img",{alt:"GoArcc Directory Structure",src:n(3141).Z})),(0,o.kt)("h2",{id:"config"},(0,o.kt)("u",null,"Config")),(0,o.kt)("p",null,"For the GoArcc application to start smoothly, we need some information. As mentioned earlier, we support gRPC, REST, GraphQL. So, we need a different port for all the servers and different settings to start the application.\nWe have a config.yaml file in the root directory. You can change any information from the file like id, password of database, etc according to your use case.\nSee the below example of the config.yaml file."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-yaml"},"\ngrpc:\n  port: 8000\n  host: localhost\n  request_timeout: 20\n\ngraphql:\n  port: 8081\n  host: localhost\n  request_Timeout: 20\n\nrest:\n  port: 8082\n  host: localhost\n  request_timeout: 20\n\n")),(0,o.kt)("h3",{id:"how-can-you-add-a-custom-config"},(0,o.kt)("u",null,"How can you add a custom config?")),(0,o.kt)("p",null,"To add some extra configs, you need to make a struct and place it inside the main struct."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'\ntype Config struct {\n    Grpc               GrpcServerConfig        `mapstructure:"GRPC"`\n    Graphql            GraphqlServerConfig     `mapstructure:"GRAPHQL"`\n    Rest               RestServerConfig        `mapstructure:"REST"`\n    HealthCheck        HealthCheckServerConfig `mapstructure:"HEALTH_CHECK"`\n    Logger             LoggerConfig            `mapstructure:"LOGGER"`\n    Postgres           PostgresConfig          `mapstructure:"POSTGRES"`\n    Metrics            MetricsConfig           `mapstructure:"METRICS"`\n    Jaeger             JaegerServerConfig      `mapstructure:"JAEGER"`\n    Auth               AuthConfig              `mapstructure:"AUTH"`\n    GithubVCSConfig    VCSSConfig              `mapstructure:"GITHUB_VCS_CONFIG"`\n    CadenceConfig      CadenceConfig           `mapstructure:"CADENCE_CONFIG"`\n}\n\n')),(0,o.kt)("div",{className:"admonition admonition-tip alert alert--success"},(0,o.kt)("div",{parentName:"div",className:"admonition-heading"},(0,o.kt)("h5",{parentName:"div"},(0,o.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,o.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"12",height:"16",viewBox:"0 0 12 16"},(0,o.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.5 0C3.48 0 1 2.19 1 5c0 .92.55 2.25 1 3 1.34 2.25 1.78 2.78 2 4v1h5v-1c.22-1.22.66-1.75 2-4 .45-.75 1-2.08 1-3 0-2.81-2.48-5-5.5-5zm3.64 7.48c-.25.44-.47.8-.67 1.11-.86 1.41-1.25 2.06-1.45 3.23-.02.05-.02.11-.02.17H5c0-.06 0-.13-.02-.17-.2-1.17-.59-1.83-1.45-3.23-.2-.31-.42-.67-.67-1.11C2.44 6.78 2 5.65 2 5c0-2.2 2.02-4 4.5-4 1.22 0 2.36.42 3.22 1.19C10.55 2.94 11 3.94 11 5c0 .66-.44 1.78-.86 2.48zM4 14h5c-.23 1.14-1.3 2-2.5 2s-2.27-.86-2.5-2z"}))),"Bright Side")),(0,o.kt)("div",{parentName:"div",className:"admonition-content"},(0,o.kt)("p",{parentName:"div"},"If you pass the environment variables then those variables will be overridden\nto the default variables which is present in the config.yaml file."))),(0,o.kt)("h3",{id:"how-you-can-pass-enviroment-variables"},(0,o.kt)("u",null,"How you can pass enviroment variables?")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},"\nGRPC_PORT = 8085\n\n")),(0,o.kt)("h2",{id:"db"},(0,o.kt)("u",null,"Db")),(0,o.kt)("p",null,"Inside the connection.go file we have a method called NewConnection(). NewConnection is responsible for opening the database connection.\nDb instance is bound with the fx provider, so we don't need to call this method.\nSee the code below."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'\nfunc NewConnection(config *config.Config) *gorm.DB {\n    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+\n        "password=%s dbname=%s sslmode=disable Timezone=Asia/Shanghai",\n        config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.Password, config.Postgres.DbName)\n\n    // Refrence taken from https://github.com/go-gorm/postgres\n    db, err := gorm.Open(postgres.New(postgres.Config{\n        DSN:                  psqlInfo,\n        PreferSimpleProtocol: true, // disables implicit prepared statement usage\n    }), &gorm.Config{})\n\n    if err != nil {\n        logger.Log.Fatal("GORM connection failed", zap.Error(err))\n        panic(err)\n    }\n\n    logger.Log.Info("connection established with the database")\n    //No need to close the connection because we have a single pool of connection.\n    return db\n}\n\n')),(0,o.kt)("p",null,"In the above code you can change the setting like timezone , ssl mode etc."),(0,o.kt)("div",{className:"admonition admonition-tip alert alert--success"},(0,o.kt)("div",{parentName:"div",className:"admonition-heading"},(0,o.kt)("h5",{parentName:"div"},(0,o.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,o.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"12",height:"16",viewBox:"0 0 12 16"},(0,o.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.5 0C3.48 0 1 2.19 1 5c0 .92.55 2.25 1 3 1.34 2.25 1.78 2.78 2 4v1h5v-1c.22-1.22.66-1.75 2-4 .45-.75 1-2.08 1-3 0-2.81-2.48-5-5.5-5zm3.64 7.48c-.25.44-.47.8-.67 1.11-.86 1.41-1.25 2.06-1.45 3.23-.02.05-.02.11-.02.17H5c0-.06 0-.13-.02-.17-.2-1.17-.59-1.83-1.45-3.23-.2-.31-.42-.67-.67-1.11C2.44 6.78 2 5.65 2 5c0-2.2 2.02-4 4.5-4 1.22 0 2.36.42 3.22 1.19C10.55 2.94 11 3.94 11 5c0 .66-.44 1.78-.86 2.48zM4 14h5c-.23 1.14-1.3 2-2.5 2s-2.27-.86-2.5-2z"}))),"Bright Side")),(0,o.kt)("div",{parentName:"div",className:"admonition-content"},(0,o.kt)("p",{parentName:"div"},"We have used GORM as an ORM. If you want to explore GORM then once go through the docs.\n(",(0,o.kt)("a",{parentName:"p",href:"https://gorm.io/docs/index.html"},"https://gorm.io/docs/index.html"),")"))),(0,o.kt)("h2",{id:"logger"},(0,o.kt)("u",null,"Logger")),(0,o.kt)("p",null,"As we know, logging is an essential part of the application. It helps in troubleshooting the application and, you can also see the performance of your infrastructure. It provides more visibility to the application at the components level. Logs contain important information so that anyone can debug or find the fault if it exists.\nWe have implemented log in the whole application so that developers don't have to worry about logging.\nIn Golang, we have multiple logging libraries. Here, we will use a zap logger."),(0,o.kt)("h3",{id:"why-have-we-used-zap-logger"},(0,o.kt)("u",null,"Why have we used zap logger? ")),(0,o.kt)("p",null,"Just see the comparison below \ud83d\ude00"),(0,o.kt)("p",null,"Log a message and 10 fields:"),(0,o.kt)("table",null,(0,o.kt)("thead",{parentName:"table"},(0,o.kt)("tr",{parentName:"thead"},(0,o.kt)("th",{parentName:"tr",align:"left"},"Package"),(0,o.kt)("th",{parentName:"tr",align:"center"},"Time"),(0,o.kt)("th",{parentName:"tr",align:"center"},"Time % to zap"),(0,o.kt)("th",{parentName:"tr",align:"center"},"Objects Allocated"))),(0,o.kt)("tbody",{parentName:"table"},(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:"left"},"\u26a1 zap"),(0,o.kt)("td",{parentName:"tr",align:"center"},"862 ns/op"),(0,o.kt)("td",{parentName:"tr",align:"center"},"+0%"),(0,o.kt)("td",{parentName:"tr",align:"center"},"5 allocs/op")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:"left"},"\u26a1 zap (sugared)"),(0,o.kt)("td",{parentName:"tr",align:"center"},"1250 ns/op"),(0,o.kt)("td",{parentName:"tr",align:"center"},"+45%"),(0,o.kt)("td",{parentName:"tr",align:"center"},"11 allocs/op")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:"left"},"zerolog"),(0,o.kt)("td",{parentName:"tr",align:"center"},"4021 ns/op"),(0,o.kt)("td",{parentName:"tr",align:"center"},"+366%"),(0,o.kt)("td",{parentName:"tr",align:"center"},"76 allocs/op")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:"left"},"go-kit"),(0,o.kt)("td",{parentName:"tr",align:"center"},"4542 ns/op"),(0,o.kt)("td",{parentName:"tr",align:"center"},"+427%"),(0,o.kt)("td",{parentName:"tr",align:"center"},"105 allocs/op")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:"left"},"apex/log"),(0,o.kt)("td",{parentName:"tr",align:"center"},"26785 ns/op"),(0,o.kt)("td",{parentName:"tr",align:"center"},"+3007%"),(0,o.kt)("td",{parentName:"tr",align:"center"},"115 allocs/op")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:"left"},"logrus"),(0,o.kt)("td",{parentName:"tr",align:"center"},"29501 ns/op"),(0,o.kt)("td",{parentName:"tr",align:"center"},"+3322%"),(0,o.kt)("td",{parentName:"tr",align:"center"},"125 allocs/op")),(0,o.kt)("tr",{parentName:"tbody"},(0,o.kt)("td",{parentName:"tr",align:"left"},"log15"),(0,o.kt)("td",{parentName:"tr",align:"center"},"29906 ns/op"),(0,o.kt)("td",{parentName:"tr",align:"center"},"+3369%"),(0,o.kt)("td",{parentName:"tr",align:"center"},"122 allocs/op")))),(0,o.kt)("p",null,"In the main method, there is a method called logger.Init(). This method will create a global logger for your application and takes log level from config.yaml file."),(0,o.kt)("p",null,"Apart from this, if we need to override the log level, we just set an env variable, or you can change the config.yaml file."),(0,o.kt)("h3",{id:"zap-log-level-reference"},(0,o.kt)("u",null,"Zap log level reference ")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"// reference taken from https://github.com/uber-go/zap/blob/120b08c8fae5be92bc7323ca78b25cd790e8c37a/level.go#L28\nconst (\n    // DebugLevel logs are typically voluminous, and are usually disabled in\n    // production.\n    DebugLevel = zapcore.DebugLevel\n    // InfoLevel is the default logging priority.\n    InfoLevel = zapcore.InfoLevel\n    // WarnLevel logs are more important than Info, but don't need individual\n    // human review.\n    WarnLevel = zapcore.WarnLevel\n    // ErrorLevel logs are high-priority. If an application is running smoothly,\n    // it shouldn't generate any error-level logs.\n    ErrorLevel = zapcore.ErrorLevel\n    // DPanicLevel logs are particularly important errors. In development the\n    // logger panics after writing the message.\n    DPanicLevel = zapcore.DPanicLevel\n    // PanicLevel logs a message, then panics.\n    PanicLevel = zapcore.PanicLevel\n    // FatalLevel logs a message, then calls os.Exit(1).\n    FatalLevel = zapcore.FatalLevel\n)\n\n")),(0,o.kt)("div",{className:"admonition admonition-tip alert alert--success"},(0,o.kt)("div",{parentName:"div",className:"admonition-heading"},(0,o.kt)("h5",{parentName:"div"},(0,o.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,o.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"12",height:"16",viewBox:"0 0 12 16"},(0,o.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.5 0C3.48 0 1 2.19 1 5c0 .92.55 2.25 1 3 1.34 2.25 1.78 2.78 2 4v1h5v-1c.22-1.22.66-1.75 2-4 .45-.75 1-2.08 1-3 0-2.81-2.48-5-5.5-5zm3.64 7.48c-.25.44-.47.8-.67 1.11-.86 1.41-1.25 2.06-1.45 3.23-.02.05-.02.11-.02.17H5c0-.06 0-.13-.02-.17-.2-1.17-.59-1.83-1.45-3.23-.2-.31-.42-.67-.67-1.11C2.44 6.78 2 5.65 2 5c0-2.2 2.02-4 4.5-4 1.22 0 2.36.42 3.22 1.19C10.55 2.94 11 3.94 11 5c0 .66-.44 1.78-.86 2.48zM4 14h5c-.23 1.14-1.3 2-2.5 2s-2.27-.86-2.5-2z"}))),"Bright Side")),(0,o.kt)("div",{parentName:"div",className:"admonition-content"},(0,o.kt)("p",{parentName:"div"},"Zap logger have different log level. You can read more about zap (",(0,o.kt)("a",{parentName:"p",href:"https://pkg.go.dev/go.uber.org/zap"},"https://pkg.go.dev/go.uber.org/zap"),")\nYou can also see the different log level of zap"))),(0,o.kt)("h2",{id:"cmd"},(0,o.kt)("u",null,"Cmd")),(0,o.kt)("p",null,"As we know, in Golang, we have a convention of cmd folders, and it has its significance. Go through this doc to read about the Go project layout (",(0,o.kt)("a",{parentName:"p",href:"https://github.com/golang-standards/project-layout"},"https://github.com/golang-standards/project-layout"),").\nLet\u2019s see our codebase. When we go inside the cmd folder, we see there are multiple files (paste GitHub link here)."),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"main.go")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"invokers.go")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"providers.go")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("strong",{parentName:"li"},"app.go"))),(0,o.kt)("div",{className:"admonition admonition-tip alert alert--success"},(0,o.kt)("div",{parentName:"div",className:"admonition-heading"},(0,o.kt)("h5",{parentName:"div"},(0,o.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,o.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"12",height:"16",viewBox:"0 0 12 16"},(0,o.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.5 0C3.48 0 1 2.19 1 5c0 .92.55 2.25 1 3 1.34 2.25 1.78 2.78 2 4v1h5v-1c.22-1.22.66-1.75 2-4 .45-.75 1-2.08 1-3 0-2.81-2.48-5-5.5-5zm3.64 7.48c-.25.44-.47.8-.67 1.11-.86 1.41-1.25 2.06-1.45 3.23-.02.05-.02.11-.02.17H5c0-.06 0-.13-.02-.17-.2-1.17-.59-1.83-1.45-3.23-.2-.31-.42-.67-.67-1.11C2.44 6.78 2 5.65 2 5c0-2.2 2.02-4 4.5-4 1.22 0 2.36.42 3.22 1.19C10.55 2.94 11 3.94 11 5c0 .66-.44 1.78-.86 2.48zM4 14h5c-.23 1.14-1.3 2-2.5 2s-2.27-.86-2.5-2z"}))),"Bright Side")),(0,o.kt)("div",{parentName:"div",className:"admonition-content"},(0,o.kt)("p",{parentName:"div"},"Before understanding the use of each file first go through dependency\nmanagement in golang.\n(",(0,o.kt)("a",{parentName:"p",href:"https://github.com/uber-go/fx"},"https://github.com/uber-go/fx"),")"))),(0,o.kt)("h3",{id:"maingo"},(0,o.kt)("u",null,"main.go")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'\npackage main\n\nimport (\n    "go.uber.org/zap"\n    "goarcc/logger"\n)\n\nfunc main() {\n    logger.Init(logger.Config{\n        LogLevel:    zap.DebugLevel, \n        Development: false,\n    })\n    GetApp().Run()\n}\n\n')),(0,o.kt)("p",null,"As we know, the main.go is the entry point for any go application. Here logger.Init() will take the config object as a parameter and initialize the whole application with debug level. This means that you will be able to see the debug levels of the log only in the standard output."),(0,o.kt)("div",{className:"admonition admonition-tip alert alert--success"},(0,o.kt)("div",{parentName:"div",className:"admonition-heading"},(0,o.kt)("h5",{parentName:"div"},(0,o.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,o.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"12",height:"16",viewBox:"0 0 12 16"},(0,o.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.5 0C3.48 0 1 2.19 1 5c0 .92.55 2.25 1 3 1.34 2.25 1.78 2.78 2 4v1h5v-1c.22-1.22.66-1.75 2-4 .45-.75 1-2.08 1-3 0-2.81-2.48-5-5.5-5zm3.64 7.48c-.25.44-.47.8-.67 1.11-.86 1.41-1.25 2.06-1.45 3.23-.02.05-.02.11-.02.17H5c0-.06 0-.13-.02-.17-.2-1.17-.59-1.83-1.45-3.23-.2-.31-.42-.67-.67-1.11C2.44 6.78 2 5.65 2 5c0-2.2 2.02-4 4.5-4 1.22 0 2.36.42 3.22 1.19C10.55 2.94 11 3.94 11 5c0 .66-.44 1.78-.86 2.48zM4 14h5c-.23 1.14-1.3 2-2.5 2s-2.27-.86-2.5-2z"}))),"Tip")),(0,o.kt)("div",{parentName:"div",className:"admonition-content"},(0,o.kt)("p",{parentName:"div"},"If you want to see the info or warn level logs then you need to just pass the Zap loglevel in config object."))),(0,o.kt)("h3",{id:"invokersgo"},(0,o.kt)("u",null,"invokers.go")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'\npackage main\n\nimport (\n    "go.uber.org/fx"\n    "goarcc/servers/cleanup"\n    "goarcc/servers/graphql"\n    "goarcc/servers/grpc"\n    "goarcc/servers/healthcheck"\n    "goarcc/servers/rest"\n)\n\n\nfunc GetInvokersOptions() fx.Option {\n    return fx.Invoke(\n        grpc.RunGRPCServer,\n        grpc.RegisterGrpcModules,\n        rest.RunRestServer,\n        graphql.RunGraphqlServer,\n        healthcheck.HealthCheckRunner,\n        cleanup.Cleanup,\n    )\n}\n\n')),(0,o.kt)("p",null,"Let's understand the above code.\nHere, we have a method GetInvokersOptions(). This function is responsible for starting multiple servers. We have gRPC, REST, and Graphql servers. Apart from that, we also have a health check server.\nCleanup code will be executed when we stop our application. Generally, in cleanup code, we gracefully shut down everything and clear the cache, if any."),(0,o.kt)("div",{className:"admonition admonition-danger alert alert--danger"},(0,o.kt)("div",{parentName:"div",className:"admonition-heading"},(0,o.kt)("h5",{parentName:"div"},(0,o.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,o.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"12",height:"16",viewBox:"0 0 12 16"},(0,o.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M5.05.31c.81 2.17.41 3.38-.52 4.31C3.55 5.67 1.98 6.45.9 7.98c-1.45 2.05-1.7 6.53 3.53 7.7-2.2-1.16-2.67-4.52-.3-6.61-.61 2.03.53 3.33 1.94 2.86 1.39-.47 2.3.53 2.27 1.67-.02.78-.31 1.44-1.13 1.81 3.42-.59 4.78-3.42 4.78-5.56 0-2.84-2.53-3.22-1.25-5.61-1.52.13-2.03 1.13-1.89 2.75.09 1.08-1.02 1.8-1.86 1.33-.67-.41-.66-1.19-.06-1.78C8.18 5.31 8.68 2.45 5.05.32L5.03.3l.02.01z"}))),"Alarm")),(0,o.kt)("div",{parentName:"div",className:"admonition-content"},(0,o.kt)("p",{parentName:"div"},"Inside Invoker, we have multiple methods. The invoker method will execute each method sequentially.\nSo there are multiple dependencies among them and if you change the order of execution you will start getting lots of errors.\nSo, if you have explored the fx library then you can play with invokers and providers."))),(0,o.kt)("h3",{id:"providersgo"},(0,o.kt)("u",null,"providers.go")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'\npackage main\n\nimport (\n    "go.uber.org/fx"\n    "goarcc/client/grpcClient"\n    "goarcc/config"\n    "goarcc/db"\n    "goarcc/servers/cleanup"\n    "goarcc/servers/grpc"\n    "goarcc/servers/openTracing/tracer/jaeger"\n)\n\nfunc GetProviderOptions() []fx.Option {\n    return []fx.Option{\n        config.ProviderFx,\n        grpc.InitGrpcBeforeServingFx,\n        db.DatabaseConnectionFx,\n        cleanup.CleanupFx,\n        jaeger.JaegerTracerFx,\n        grpcClient.ConnectionFx,\n    }\n}\n\n')),(0,o.kt)("p",null,"Let's understand the above code.\nHere We have a method GetProvidersOptions(). This function loads all the dependencies which we will need to start the application. There are multiple servers and each server has some dependency or needs some predefined object at runtime."),(0,o.kt)("h3",{id:"appgo"},(0,o.kt)("u",null,"app.go")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'\npackage main\n\nimport "go.uber.org/fx"\n\nfunc GetApp() *fx.App {\n    opts := make([]fx.Option, 0)\n    opts = GetProviderOptions()\n    opts = append(opts, GetInvokersOptions())\n    return fx.New(\n        opts...,\n    )\n}\n\n')),(0,o.kt)("p",null,"Here the GetApp() method will resolve all the dependencies and bind the whole application as an object. So, you can easily start the application in the main file."),(0,o.kt)("h2",{id:"servers"},(0,o.kt)("u",null," Servers ")),(0,o.kt)("p",null,(0,o.kt)("img",{alt:"GoArcc Servers Structure",src:n(5714).Z})),(0,o.kt)("p",null,"Servers are the core part of this project. As we mentioned, we support multiple servers. Let's dive deep into the codebase."),(0,o.kt)("h3",{id:"grpc"},"GRPC"),(0,o.kt)("p",null,"GoArcc is built upon client-server architecture. So, GRPC is the core server for GoArcc. Every request is handled by the GRPC server and middlewares."),(0,o.kt)("h4",{id:"grpc-middleware"},"GRPC Middleware"),(0,o.kt)("p",null,"You have to set up your middleware into gRPC, and it will work for both REST and GraphQL. Auth middleware is already implemented here."),(0,o.kt)("h4",{id:"how-to-register-your-service-client"},"how to register your service client?"),(0,o.kt)("p",null,"you can register your service as gRPC at ",(0,o.kt)("inlineCode",{parentName:"p"},"servers/grpc/register.go")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"\nfunc RegisterGrpcModules(srv *grpc.Server,\n    db *gorm.DB,\n    config *config.Config,\n    grpcClientConnection *grpc.ClientConn) {\n\n    //register new grpc modules here\n    // RegisterUserProfilesServer autogenerated in pb/..._grpc.go file\n    userProfilePb.RegisterUserProfilesServer(srv, userExt.NewUserProfilesServer(db, config, grpcClientConnection))\n}\n\n")),(0,o.kt)("h3",{id:"rest"},"REST"),(0,o.kt)("p",null,"REST APIs over gRPC are achieved by ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/grpc-ecosystem/grpc-gateway"},"gRPC-gateway"),". It translates a RESTful HTTP API into gRPC."),(0,o.kt)("h4",{id:"how-to-register-your-service-as-rest"},"how to register your service as REST?"),(0,o.kt)("p",null,"you can register your service as rest at ",(0,o.kt)("inlineCode",{parentName:"p"},"servers/rest/register.go")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'func RegisterRESTModules(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {\n   // Register your REST module here\n   // RegisterUserProfilesHandler autogenerated in pb/..._gw.go\n    if err := userProfilePb.RegisterUserProfilesHandler(ctx, mux, conn); err != nil {\n        logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))\n        return err\n    }\n    return nil\n}\n')),(0,o.kt)("h3",{id:"graphql"},"GraphQL"),(0,o.kt)("p",null,"We are using ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/ysugimoto/grpc-graphql-gateway"},"grpc-graphql-gateway"),", this plugin generates GraphQL Schema from Protocol Buffers."),(0,o.kt)("h4",{id:"how-to-register-your-service-as-graphql"},"how to register your service as graphQL?"),(0,o.kt)("p",null,"you can register your service as rest at ",(0,o.kt)("inlineCode",{parentName:"p"},"servers/graphql/register.go")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'//RegisterGraphqlModules: Mapping the services with the single graphql endpoint\nfunc RegisterGraphqlModules(mux *runtime.ServeMux, conn *grpc.ClientConn) error {\n    // Register your graphql service here\n    if err := userProfilePb.RegisterUserProfilesGraphqlHandler(mux, conn); err != nil {\n        logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))\n        return err\n    }\n    return nil\n}\n')),(0,o.kt)("h3",{id:"healthcheck"},"HealthCheck"),(0,o.kt)("p",null,"Any client will be able to check the health of our product. Health information provides a brief idea to the user if the service of their project is down.\nHealthCheck runs on ",(0,o.kt)("inlineCode",{parentName:"p"},"http://localhost:8083/health/")),(0,o.kt)("h3",{id:"jaeger"},"Jaeger"),(0,o.kt)("p",null,"Jaeger is open-source software for tracing transactions between distributed services. It\u2019s used for monitoring and troubleshooting complex microservices environments.\nOpen http://loaclhost:16686 to trace each request easily."),(0,o.kt)("h3",{id:"prometheus"},"Prometheus"),(0,o.kt)("p",null,"Prometheus is an open-source system for monitoring and alerting.\nOpen ",(0,o.kt)("inlineCode",{parentName:"p"},"http://localhost:9090"),"."))}g.isMDXComponent=!0},5714:function(e,t,n){"use strict";t.Z=n.p+"assets/images/Servers-4f5ba225c2c6fa9c61844e7ec6e2b440.jpg"},3141:function(e,t,n){"use strict";t.Z=n.p+"assets/images/goArcc-bb2276616994c26a9456fb077125f209.jpg"}}]);