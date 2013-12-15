# Points

![screenshot](https://raw.github.com/jeffcarp/points/master/screenshot.png)

A simple one-page app written in Go that lets you award points to people in your office. Start by filling in people's information in **people.json**. Run with:

```bash
go run app.go
```

Then head over to [localhost:9000](http://localhost:9000/).

To access the admin panel, take a secret word and MD5 it once. This is your pass key. Then, take your pass key and MD5 it again. Take the double-hashed key and update the 2 parts of the source code that check against it. You can access the admin panel by going to the index page with your 1st key as a GET parameter, e.g. localhost:9000/?key=the_first_md5_you_created
