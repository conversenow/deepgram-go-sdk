// Copyright 2023-2024 Deepgram SDK contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	prettyjson "github.com/hokaccha/go-prettyjson"

	prerecorded "github.com/deepgram/deepgram-go-sdk/pkg/api/prerecorded/v1"
	interfaces "github.com/deepgram/deepgram-go-sdk/pkg/client/interfaces"
	client "github.com/deepgram/deepgram-go-sdk/pkg/client/prerecorded"
)

const (
	filePath string = "./Bueller-Life-moves-pretty-fast.mp3"
)

func main() {
	// init library
	client.Init(client.InitLib{
		LogLevel: client.LogLevelDefault, // LogLevelStandard / LogLevelFull / LogLevelTrace
	})

	// context
	ctx := context.Background()

	// set the Transcription options
	options := &interfaces.PreRecordedTranscriptionOptions{
		Punctuate:  true,
		Diarize:    true,
		Language:   "en-US",
		Utterances: true,
	}

	//client
	c := client.NewWithDefaults()
	dg := prerecorded.New(c)

	// open file sream
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("os.Open(%s) failed. Err: %v\n", filePath, err)
		os.Exit(1)
	}
	defer file.Close()

	// send stream to Deepgram
	res, err := dg.FromStream(ctx, file, options)
	if err != nil {
		fmt.Printf("FromStream failed. Err: %v\n", err)
		os.Exit(1)
	}

	data, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("json.Marshal failed. Err: %v\n", err)
		os.Exit(1)
	}

	// make the JSON pretty
	prettyJSON, err := prettyjson.Format(data)
	if err != nil {
		fmt.Printf("prettyjson.Marshal failed. Err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n\nResult:\n%s\n\n", prettyJSON)
}
