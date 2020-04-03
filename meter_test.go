<<<<<<< HEAD
/**
 * Copyright (c) 2019 eBay Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 **/
=======
>>>>>>> 20be12848571a13b5c53c7a664f21a70beee615a
package goovn

import (
	"testing"
)

<<<<<<< HEAD
func TestMeter(t *testing.T) {
	var cmds []*OvnCommand
	cmd, err := ovndbapi.MeterAdd("meter1", "drop", 101, "kbps", nil, 300)
	if err != nil {
		t.Fatal(err)
	}
	cmds = append(cmds,cmd)
	cmd, err = ovndbapi.MeterAdd("meter2", "drop", 101, "kbps", nil, 300)
	if err != nil {
		t.Fatal(err)
	}
	cmds = append(cmds,cmd)
	cmd, err = ovndbapi.MeterAdd("meter3", "drop", 101, "kbps", nil, 300)
	if err != nil {
		t.Fatal(err)
	}
	cmds = append(cmds,cmd)
=======
const (
	METER1 = "testMeter1"
	METER2 = "testMeter2"
	METER3 = "testMeter3"
)

func TestMeter(t *testing.T) {
	ovndbapi := getOVNClient(DBNB)
	var cmds []*OvnCommand
	cmd, err := ovndbapi.MeterAdd(METER1, "drop", 101, "kbps", nil, 300)
	if err != nil {
		t.Fatal(err)
	}
	cmds = append(cmds, cmd)
	cmd, err = ovndbapi.MeterAdd(METER2, "drop", 101, "kbps", nil, 300)
	if err != nil {
		t.Fatal(err)
	}
	cmds = append(cmds, cmd)
	cmd, err = ovndbapi.MeterAdd(METER3, "drop", 101, "kbps", nil, 300)
	if err != nil {
		t.Fatal(err)
	}
	cmds = append(cmds, cmd)
>>>>>>> 20be12848571a13b5c53c7a664f21a70beee615a
	err = ovndbapi.Execute(cmds...)
	if err != nil {
		t.Fatal(err)
	}

<<<<<<< HEAD

	meter,err := ovndbapi.MeterList()
	if err != nil{
		t.Fatal(err)
	}
	if len(meter)!=3{
=======
	meter, err := ovndbapi.MeterList()
	if err != nil {
		t.Fatal(err)
	}
	if len(meter) < 3 {
>>>>>>> 20be12848571a13b5c53c7a664f21a70beee615a
		t.Fatal("Meter add Fail")
	}

	meterBands, err := ovndbapi.MeterBandsList()
<<<<<<< HEAD
	if err != nil{
		t.Fatal(err)
	}
	if len(meterBands)!=3{
		t.Fatal("Meter bands shows Fail")
	}

	cmd ,err = ovndbapi.MeterDel("meter1")
	if err != nil {
		t.Fatal(err)
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		t.Fatal(err)
	}
	meter,err = ovndbapi.MeterList()
	if err != nil{
		t.Fatal(err)
	}
	if len(meter) != 2{
		t.Fatal("Delete single Meter Error")
	}

	cmd ,err = ovndbapi.MeterDel()
	if err != nil {
		t.Fatal(err)
	}
	err = ovndbapi.Execute(cmd)
	if err != nil {
		t.Fatal(err)
	}
	meter,err = ovndbapi.MeterList()
	if len(meter)!=0{
		t.Fatal("Delete All Meter Fail")
	}

}
=======
	if err != nil {
		t.Fatal(err)
	}
	if len(meterBands) < 3 {
		t.Fatal("Meter bands shows Fail")
	}

	defer func() {
		cmd, err = ovndbapi.MeterDel()
		if err != nil {
			t.Fatal(err)
		}
		err = ovndbapi.Execute(cmd)
		if err != nil {
			t.Fatal(err)
		}
		meter, err = ovndbapi.MeterList()
		if err != nil {
			t.Fatal(err)
		}
		if len(meter) != 0 {
			t.Fatal("Delete All Meter Fail")
		}
	}()

	defer func() {
		cmd, err = ovndbapi.MeterDel(METER1)
		if err != nil {
			t.Fatal(err)
		}
		err = ovndbapi.Execute(cmd)
		if err != nil {
			t.Fatal(err)
		}
		meter, err = ovndbapi.MeterList()
		if err != nil {
			t.Fatal(err)
		}
		if len(meter) < 2 {
			t.Fatal("Delete single Meter Error")
		}
	}()
}
>>>>>>> 20be12848571a13b5c53c7a664f21a70beee615a
