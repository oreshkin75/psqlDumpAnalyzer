// +build macos

package platform

func (d *DumpCreator) CreateDump() ([]string, error) {
	d.logger.LogInfo.Printf("in macos you can't create dump, but can analyze your windows dumps")
	return nil, nil
}
