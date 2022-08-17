# UV

## Introduction

UV is a command line application written in Go that checks the current level of UV radiation expressed on the UV Index scale. Additionally, it displays a set of times related to the Sun's position in the sky and an estimation of safe sunlight exposure for different skin types.

## Setup

Binary files are released from time to time. However, users are encouraged to build the application on their own. The only requirement is a Go compiler and building scripts are provided for Linux and Windows operating systems. The application's directory should be appended to PATH environmental variable. 

For the core functionality, [OpenUV](https://www.openuv.io/) API key is required. If the user wants to specify location by city and country instead of latitude and longitude, [OpenWeather](https://openweathermap.org/) API key is needed as well.

### Settings

Settings are stored in a file under `./settings/uv.json`. If the file is not found, a new one is generated. With the help of this file the user provides their API keys and, optionally, a default location of choice (may be left as `null`). Specifying a default location allows to omit the passage of the same name over and over again. The default location may also be set with an appropriate flag. The request limit allows to track the number of daily requests left on OpenUV API account. Setting it to `-1` disables tracking.

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
    },

    "request_limit": 50
}
```

Note that southern latitude and western longitude are negative.

## Usage

### Flags

| Flag | Description                                                             |
|------|-------------------------------------------------------------------------|
| `-d` | make the location specified with the `-l` flag the new default location |
| `-h` | display help                                                            |
| `-l` | specify location, otherwise the default location will be used           |
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

## Output

### Structure

The description of most elements is based on the [OpenUV API Documentation](https://www.openuv.io/uvindex). Time is expressed in UTC.

| Element             | Description                                                               |
|---------------------|---------------------------------------------------------------------------|
| Header              | Location name and coordinates, generation time, [requests left]           |
| UV Index block      | UV index - current, daily maximum (value and time) and ozone level \[du\] |
| Sunrise             | Sunrise time                                                              |
| Solar Noon          | Sun's zenith time                                                         |
| Sunset              | Sunset time                                                               |
| Night               | Night becomes dark enough for astronomical observations                   |
| Golden Hour         | Evening golden hour's start time                                          |
| Morning GH ends     | Morning golden hour's end time                                            |
| Safe Exposure block | Safe sunlight exposure time for skin types on Fitzpatrick scale           |

### Example 

This is the example report for the city of London.

```
London, GB    51.507 -0.128    2022-08-17 07:43 UTC    [43]

UV Index:
  Current:   1.27
  Max:       5.59 (12:06)
  Ozone:   332.50

Sunrise:           04:50
Solar Noon:        12:06
Sunset:            19:21
Night:             21:43
Golden Hour:       18:35
Morning GH ends:   05:36

Safe Exposure Time [min]:
  1:   131   |   4:   262
  2:   157   |   5:   420
  3:   210   |   6:   787
```

## Resources

* [UV Index](https://en.wikipedia.org/wiki/Ultraviolet_index)
* [Fizpatrick scale](https://en.wikipedia.org/wiki/Fitzpatrick_scale)

* [OpenUV](https://www.openuv.io/)
* [OpenWeather](https://openweathermap.org/)

## License

This software is available under MIT License.
