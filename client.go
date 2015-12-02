package userstream

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	userStreamsEndPoint = "https://userstream.twitter.com/1.1/user.json"
	requestTokenUrl     = "https://api.twitter.com/oauth/request_token"
	authorizeTokenUrl   = "https://api.twitter.com/oauth/authorize"
	accessTokenUrl      = "https://api.twitter.com/oauth/access_token"

	twitterApiUrl = "https://api.twitter.com"

	requestTokenPath   = "/oauth/request_token"
	authorizeTokenPath = "/oauth/authorize"
	accessTokenPath    = "/oauth/access_token"
)

type Client struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
	cons              *oauth.Consumer
	token             *oauth.AccessToken
}

func (c *Client) UserStream(callback func(interface{})) {
	c.connect(userStreamsEndPoint, callback)
}

func (c *Client) FollowUserId(tweetId int64) (*UserDetails, error) {
	user := UserDetails{}
	id := fmt.Sprintf("%d", tweetId)
	url := fmt.Sprintf("/1.1/friendships/create.json")
	fmt.Println(url)
	fmt.Println(id)
	response, err := c.post(
		c.apiUrl(url),
		map[string]string{
			"user_id": "1401881",
			"follow":  "true",
		},
	)
	if err != nil {
		return &user, err
	}

	return c.userDetailsByResponse(response)
}

func (c *Client) connect(endPoint string, callback func(interface{})) {
	consumer := oauth.NewConsumer(
		c.ConsumerKey,
		c.ConsumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   requestTokenUrl,
			AuthorizeTokenUrl: authorizeTokenUrl,
			AccessTokenUrl:    accessTokenUrl,
		},
	)

	response, err := consumer.Post(endPoint, nil, c.accessToken())
	if err != nil {
		log.Fatal(err)
	}

	c.readStream(response, callback)
}

func (c *Client) readStream(response *http.Response, callback func(interface{})) {
	reader := bufio.NewReader(response.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}

		if len(line) > 0 && string(line) != "\r\n" {
			object := ParseJson(string(line))
			callback(object)
		}
	}
}

func (c *Client) accessToken() *oauth.AccessToken {
	return &oauth.AccessToken{
		Token:  c.AccessToken,
		Secret: c.AccessTokenSecret,
	}

}

func (c *Client) get(requestUrl string, params map[string]string) (*http.Response, error) {
	return c.consumer().Get(requestUrl, params, c.accessToken())
}

func (c *Client) post(requestUrl string, params map[string]string) (*http.Response, error) {
	return c.consumer().Post(requestUrl, params, c.accessToken())
}

func (c *Client) consumer() *oauth.Consumer {
	if c.cons != nil {
		return c.cons
	}

	c.cons = oauth.NewConsumer(
		c.ConsumerKey,
		c.ConsumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   c.apiUrl(requestTokenPath),
			AuthorizeTokenUrl: c.apiUrl(authorizeTokenPath),
			AccessTokenUrl:    c.apiUrl(accessTokenPath),
		},
	)
	return c.cons
}

func (c *Client) apiUrl(format string, a ...interface{}) string {
	apiPath := fmt.Sprintf(format, a...)
	return twitterApiUrl + apiPath
}

func (c *Client) userDetailsByResponse(response *http.Response) (*UserDetails, error) {
	decoder := c.jsonDecoder(response)
	user := UserDetails{}
	decoder.Decode(&user)
	spew.Dump(user)
	return &user, nil
}

func (c *Client) tweetsByResponse(response *http.Response) ([]Tweet, error) {
	decoder := c.jsonDecoder(response)
	tweets := []Tweet{}
	decoder.Decode(&tweets)
	return tweets, nil
}

func (c *Client) jsonDecoder(response *http.Response) *json.Decoder {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return json.NewDecoder(bytes.NewReader(data))
}
