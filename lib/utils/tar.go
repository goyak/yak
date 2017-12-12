// Copyright © 2017 EasyStack Inc. Shawn Wang <shawn.wang@easystack.cn>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package utils

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func AddTarFile(tw *tar.Writer, path string) error {
	fmt.Printf("prepare: %s.\n", path)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if stat, err := file.Stat(); err == nil {
		// now lets create the header as needed for this file within the tarball
		header := new(tar.Header)
		header.Name = path
		if !stat.IsDir() {
			header.Size = stat.Size()
			header.ModTime = stat.ModTime()
			header.Mode = int64(stat.Mode())
		}

		if err := tw.WriteHeader(header); err != nil {
			fmt.Printf(" >> n: %s \n", header.Name)
			fmt.Printf(" >> size: %d \n", header.Size)
			fmt.Printf(" >> m: %d \n", header.Mode)
			fmt.Printf(" >> err: %s \n", err)
			return err
		}

		if !stat.IsDir() {
			// write the header to the tarball archive
			// copy the file data to the tarball
			fmt.Printf("copying %s ... m: %d\n", path, header.Mode)
			if _, err := io.Copy(tw, file); err != nil {
				return err
			}
			fmt.Printf("copied %s.\n", path)
		}
	}
	return nil
}
