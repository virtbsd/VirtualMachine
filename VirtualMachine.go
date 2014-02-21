/*
(BSD 2-clause license)

Copyright (c) 2014, Shawn Webb
All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

   * Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package VirtualMachine

import (
    "github.com/coopernurse/gorp"
)

type VirtualMachine interface {
    GetUUID() string
    Start() error
    Stop() error
    Status() string
    CreateSnapshot(snapname string) error
    RestoreSnapshot(snapname string) error
    DeleteSnapshot(snapname string) error
    PrepareHostNetworking() error
    PrepareGuestNetworking() error
    NetworkingStatus() string
    GetPath() (string, error)
    IsOnline() bool
    Validate() error
    Persist(db *gorp.DbMap) error
    Delete(db *gorp.DbMap) error
    Archive(archivename string) error
}

type VirtualMachineError struct {
    ErrorString string
    VM VirtualMachine
}

func (e VirtualMachineError) Error() string {
    return "[" + e.VM.GetUUID() + "]: " + e.ErrorString
}
