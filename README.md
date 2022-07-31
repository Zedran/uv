# UV

## Introduction

UV is an app written in Go that checks the current level of UV radiation measured on the UV Index scale.

## Setup

For the core functionality, [OpenUV](https://www.openuv.io/) API key is required. If the user wants to specify location by city and country instead of latitude and longitude, [OpenWeather](https://openweathermap.org/) API key is needed as well.

### Settings

Settings are stored in a file under `./settings/uv.json`. If the file is not found, a new one is generated. With the help of this file the user provides their API keys and, optionally, a default location of choice (may be left as `null`). Specifying a default location allows to omit the passage of the same coordinates over and over again.

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

## Resources

* [UV Index](https://en.wikipedia.org/wiki/Ultraviolet_index)

* [OpenUV](https://www.openuv.io/)
* [OpenWeather](https://openweathermap.org/)

## License

This software is available under MIT License.
