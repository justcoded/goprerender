# Change Log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

-------------------------------------------------------------------------

## [1.2.0]() - _2021-07-01_

### Added

* Added support to pass `CACHE_EXPIRATION_TIME` environment variable

### Changed

* Only the `Clear-Site-Data: *` header purges all cache now
* The `Cache-Control: must-revalidate` header removes only one page instead of flushing all caches

-------------------------------------------------------------------------

## [1.1.0]() - _2022-05-13_

### Added

* Added an ability to purge all cache by sending `Cache-Control: must-revalidate` or `Clear-Site-Data: *` header to the
  prerender request

-------------------------------------------------------------------------

## [1.0.0]() - _2021-10-05_

### Added

* Taken https://github.com/goprerender/prerender as a base
* Added support to pass `STORAGE_URL` environment variable
* Dockerized goprerender server and storage
* Added docker-compose example and Makefile commands for faster setup

-------------------------------------------------------------------------
