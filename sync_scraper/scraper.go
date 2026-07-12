package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/html"
)

var domain string;

func getLinks(n *html.Node, links []string) []string {
	// filter out non links
	if (n.Type == html.ElementNode && n.Data == "a") {
		// search for 'href' to find link	
		for _, a := range n.Attr {
			if (a.Key == "href") {
				links = append(links, a.Val);
				break;
			}
		}	
	} 
	
	// search other children
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = getLinks(c, links);
	}

	return links;
}

func generateUrl(link string) string {
	// check if first character is '/'	
	if (link[0] == 47) {
		return domain + link;
	} else {
		return link;
	}
}

func shouldRecurse(url string) bool {
	i := min(len(domain), len(url));
	if (url[:i] == domain) {
		return true;
	} else {
		return false;
	}
}

func visitUrl(url string, visited, failed []string) ([]string, []string) {
	// skip if visited	
	if (isVisited(url, visited)) {
		return visited, failed;
	} else {
		visited = append(visited, url);
	}
	
	fmt.Println(url);
	res, err := http.Get(url);
	if (err != nil) {
		fmt.Println("Error:", err);
		failed = append(failed, url);
		return visited, failed;
	}
	defer res.Body.Close();

	doc, err := html.Parse(res.Body);
	if (err != nil) {
		fmt.Println("Error:", err);
		return visited, failed; 
	}
	
	var links []string;
	fmt.Printf("Url: %s, Status Code: %d\n", url, res.StatusCode);

	if (res.StatusCode == 404) {
		failed = append(failed, url);
	}

	if (res.StatusCode == 200) {
		links = getLinks(doc, links);
	}
	
	if (shouldRecurse(url)) {
		for _, link := range links {
			visited, failed = visitUrl(generateUrl(link), visited, failed);
		}
	}

	return visited, failed;
}

func isVisited(url string, visited []string) bool {
	for _, link := range visited {
		if (url == link) {
			return true;
		}
	}

	return false;
}

func main() {
	requestUrl := "http://localhost:8080/";
	domain = "http://localhost:8080/";
	
	var visited []string;
	var failed []string;
	visited, failed = visitUrl(requestUrl, visited, failed);

	for _, link := range failed {
		fmt.Println("Failed: ", link);
	}
}
