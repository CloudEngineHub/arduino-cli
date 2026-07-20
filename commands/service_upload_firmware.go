// This file is part of arduino-cli.
//
// Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

package commands

/*

// UploadFirmwareFileToServerStreams return a server stream that forwards the output and error streams to the provided writers.
// It also returns a function that can be used to retrieve the result of the upload.
func UploadFirmwareFileToServerStreams(ctx context.Context, outStream io.Writer, errStream io.Writer) (rpc.ArduinoCoreService_UploadFirmwareFileServer, func() *rpc.UploadResult) {
	var result *rpc.UploadResult
	stream := streamResponseToCallback(ctx, func(resp *rpc.UploadFirmwareFileResponse) error {
		if errData := resp.GetErrStream(); len(errData) > 0 {
			_, err := errStream.Write(errData)
			return err
		}
		if outData := resp.GetOutStream(); len(outData) > 0 {
			_, err := outStream.Write(outData)
			return err
		}
		if res := resp.GetResult(); res != nil {
			result = res
		}
		return nil
	})
	return stream, func() *rpc.UploadResult {
		return result
	}
}

// UploadFirmwareFile performs the upload of a firmware file to a board.
func (s *arduinoCoreServerImpl) UploadFirmwareFile(req *rpc.UploadFirmwareFileRequest, stream rpc.ArduinoCoreService_UploadFirmwareFileServer) error {
	return fmt.Errorf("not implemented")
}
*/
