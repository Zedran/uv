# UV

## Introduction

UV is a command line application written in Go that checks the current level of UV radiation measured on the UV Index scale.

## Setup

For the core functionality, [OpenUV](https://www.openuv.io/) API key is required. If the user wants to specify location by city and country instead of latitude and longitude, [OpenWeather](https://openweathermap.org/) API key is needed as well.

### Settings

Settings are stored in a file under `./settings/uv.json`. If the file is not found, a new one is generated. With the help of this file the user provides their API keys and, optionally, a default location of choice (may be left as `null`). Specifying a default location allows to omit the passage of the same name over and over again. The default location may also be specified with an appropriate flag.

The settings file is structured as follows:

```json
{
    "open_weather_key": "key1",
    "open_uv_key":      "key2",

    "default_location": {
        "name":    "London",
        "country": "GB",
        "lat":     51.50732,
        "lon":     -0.1276474
    }
}
```

Note that southern latitude and western longitude are negative.

## Usage

### Flags

| Flag | Description                                                             |
|------|-------------------------------------------------------------------------|
| `-l` | specify location, otherwise the default location will be used           |
| `-d` | make the location specified with the `-l` flag the new default location |
| `-u` | unset the current default location                                      |

If the city name has more than one word in it, wrap it in quotes. If you want to specify a country name as well, separate it from the city name with a comma and put the whole thing in quotes.

### Examples

| Syntax                   | Meaning                                            |
|--------------------------|----------------------------------------------------|
| `uv -l London`           | query the API with city name only as the location  |
| `uv -l "Rio de Janeiro"` | query the API with a multi-word city name          |
| `uv -l "London, GB"`     | query the API by specifying city and country       |
| `uv -l London -d`        | query the API and make London the default location |
| `uv`                     | query the API using the default location           |
| `uv -u`                  | unset the current default location                 |

## Resources

* [UV Index](https://en.wikipedia.org/wiki/Ultraviolet_index)
* [Fizpatrick scale](https://en.wikipedia.org/wiki/Fitzpatrick_scale)

* [OpenUV](https://www.openuv.io/)
* [OpenWeather](https://openweathermap.org/)

## License

This software is available under MIT License.
