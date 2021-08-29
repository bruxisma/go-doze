// Copyright 2021 Isabella Muerte. All rights reserved.
//
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE.md file that accompanied this package

/*
Doze is a library that wraps the resty golang library to more closely focus its
API for us in writing small and quick REST wrappers for use in terraform
provider plugins. Additionally, it provides a test module to make testing the
REST APIs you write easier when a given platform does *not* provide a sandbox
API to test against

Usage

The doze package is designed to be used with go modules enabled only.

	import "occult.work/doze"
	import "occult.work/doze/test"

Examples:

Construct a Client, then get a Request object to perform requests with.

	client := doze.NewClient().
		SetTimeout(time.Duration(1 * time.Minute)).
		SetAuthScheme("Basic").
		SetAuthToken(token).
		SetHostURL(BaseURL).
		SetError(&Error{})
	request = client.Request(ctx, &ReturnType{})

*/
package doze
