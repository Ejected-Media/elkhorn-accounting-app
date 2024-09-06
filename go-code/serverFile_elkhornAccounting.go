// . . Elkhorn - Accounting - Website App

package main
// .

import (
  "os"
  "log"
		
  "text/template"
  "net/http"
		
  "cloud.google.com/go"
		
  "cloud.google.com/go/firestore"
  "google.golang.org/api/iterator"
)



type htmlPageData struct {
    pageTitle string
    pagePath string
    pageList []pageNav
    
}


type pageNav struct {
    pageTitle string
    pageLink string
}

// . indexHandler,  ~ for Public Pages °
func indexHandler(w http.ResponseWriter, r *http.Request) {
// ,  ° . +
    if r.URL.Path != "/" {
    	http.NotFound(w, r)
    	return
    }

// , ° . +
  pageTitle := "~ . Elkhorn Accounting App - // - Website App"
  pagePath := r.URL.Path
  // pageType := ".."


// ,  ° . +
pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
    //  pageList: []pageNav {
  //        { pageTitle: "one", pageLink: "one"},
//          { pageTitle: "two", pageLink: "two"},
   //       { pageTitle: "three", pageLink: "three"},
    //  },
  	
  }  //. .  pageData
  
  
  if pagePath == "/" {
      pageTitle = "Index Page"
 //     pageList = pageList
  }


// ,  ° . +
  pageFilePath := template.Must(template.ParseFiles("layout_main_page.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  indexHandler


//  .  html url routes 
//  .  as well as terminal cli logs
func main() {
// ,  ° . +
  appName := "~ . Elkhorn Accounting App - // - Website App"

// ,  ° . +
  http.HandleFunc("/", indexHandler)

// . ° ~ +
 http.HandleFunc("/action/classSchedule", actionPage)

// -- -
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
    log.Printf("Loading _webapp with default port")
  }
  
// ,  ° . +
  log.Printf("_webapp is active and Listening on port %s", port)

  log.Printf("// -- - %s", appName)
  log.Printf("_webapp now loaded and running at http://localhost:%s", port)

// -- - 
  if err := http.ListenAndServe(":"+port, nil); err != nil {
    log.Fatal("Error Starting the HTTP Server :", err)
    return
  }}

