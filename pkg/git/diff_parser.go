	unrecognized   parserState = "unrecognized"
	diffBegin      parserState = "diffBegin"
	diffBody       parserState = "diffBody"
	newFileDiff    parserState = "newFileDiff"
	deleteFileDiff parserState = "deleteFileDiff"
	modifyFileDiff parserState = "modifyFileDiff"
	ignoreDiff     parserState = "ignoreDiff"
		if strings.HasPrefix(line, "deleted file ") {
		if strings.HasPrefix(line, "new file ") {
		return fmt.Errorf("unexpected diff line: %#v", line)
	pathA := strings.TrimPrefix(lineParts[2], "a/")
	pathB := strings.TrimPrefix(lineParts[3], "b/")
	if !p.PathFilter.IsFilePathValid(pathA) || !p.PathFilter.IsFilePathValid(pathB) {
		p.state = ignoreDiff
		return nil
	newPathA := p.PathFilter.TrimFileBasePath(pathA)
	newPathB := p.PathFilter.TrimFileBasePath(pathB)
	newLine := fmt.Sprintf("diff --git a/%s b/%s", newPathA, newPathB)