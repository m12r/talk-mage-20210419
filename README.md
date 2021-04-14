# Talk: Mage 2021-04-19

A talk about [Mage][] held at [GoGraz April 2021 Meetup][go-graz-202104].

## Required Software

- go >= v1.16
- node >= v14
- xdg-open

**Note:** if `xdg-open` is not installed, it will write an error to `stderr` 
and you have to manually open your web browser.

## Usage

Just run the following command. It starts a http server, and a web browser
where the presentation is shown.

```shell
go run main.go
```

The example code uses mage for building.

```shell
cd code
mage
```

**Note:** instead of `mage`, you can also run `go run mage`, if you don't have
`mage` installed.

## License

This project is licensed under the Creative Commons BY-CA License version 4.0.

<a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/"><img alt="Creative Commons License BY-CA version 4.0" style="border-width:0" src="https://i.creativecommons.org/l/by-sa/4.0/88x31.png" /></a>

## Dependencies

- [Go][] 
- [Mage][] mascot »Gary«
- [Remark][]

---

© 2021, [Matthias Endler][me].


[go]: https://go.dev
[go-graz-202104]: https://gograz.org/meetup/2021-04-19/
[mage]: https://magefile.org
[me]: https://m12r.at
[remark]: https://remarkjs.com