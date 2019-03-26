// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package topic

type Topic struct {
	Name              string
	NumPartitions     int32
	ReplicationFactor int16
}

func (t Topic) Equals(topic Topic) bool {
	if t.Name != topic.Name {
		return false
	}
	if t.NumPartitions != topic.NumPartitions {
		return false
	}
	if t.ReplicationFactor != topic.ReplicationFactor {
		return false
	}
	return true
}
