package goutil

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// RangeDir 遍历指定文件夹下的所有子目录和文件，返回文件夹和文件地址。
func RangeDir(dir string) ([]string, error) {
	var temp []string
	return getFileAll(dir, temp)
}

// getFileAll 遍历指定文件夹的目录结构
// 参数path：要遍历的文件夹，注意，末尾不要加斜杠，例如：C:
// 参数files：用于临时存储文件夹或文件信息
// 返回值1：文件及文件夹的地址信息
// 返回值2：错误
func getFileAll(dirPath string, files []string) ([]string, error) {
	fInfo, err := os.ReadDir(dirPath)
	if err != nil {
		return files, err
	}
	for _, fi := range fInfo {
		if fi.IsDir() {
			tempDir := filepath.Join(dirPath, fi.Name())
			files = append(files, tempDir)
			files, _ = getFileAll(tempDir, files)
		} else {
			files = append(files, filepath.Join(dirPath, fi.Name()))
		}
	}
	return files, nil
}

// RangeDirFile 遍历指定文件夹(含所有子目录)下的所有文件，仅返回文件地址。
func RangeDirFile(dir string) ([]string, error) {
	var temp []string
	return getFileAllFile(dir, temp)
}

// getFileAll 遍历指定文件夹的目录结构
// 参数path：要遍历的文件夹，注意，末尾不要加斜杠，例如：C:
// 参数files：用于临时存储文件夹或文件信息
// 返回值1：文件及文件夹的地址信息
// 返回值2：错误
func getFileAllFile(dirPath string, files []string) ([]string, error) {
	fInfo, err := os.ReadDir(dirPath)
	if err != nil {
		return files, err
	}
	for _, fi := range fInfo {
		if fi.IsDir() {
			tempDir := filepath.Join(dirPath, fi.Name())
			//files = append(files, tempDir)
			files, _ = getFileAllFile(tempDir, files)
		} else {
			files = append(files, filepath.Join(dirPath, fi.Name()))
		}
	}
	return files, nil
}

// RangeDirFileFilter 遍历指定文件夹(含所有子目录)下的所有文件，仅返回文件地址。
func RangeDirFileFilter(dir string, filterFunc func(filename string) bool) ([]string, error) {
	var temp []string
	return getFileAllFileFilter(dir, temp, filterFunc)
}

// FileFilterByType 返回一个根据文件扩展名进行筛选的闭包函数
// ext：文件扩展名，比如 .txt .mp4
func FileFilterByType(ext string) func(filename string) bool {
	ext = strings.ToLower(ext)
	return func(filename string) bool {
		return strings.ToLower(filepath.Ext(filename)) == ext
	}
}

// FileFilterByTypes 返回一个根据文件扩展名进行筛选的闭包函数
// ext：文件扩展名列表，比如 .txt .mp4
func FileFilterByTypes(ext []string) func(filename string) bool {
	return func(filename string) bool {
		_ext := strings.ToLower(filepath.Ext(filename))
		for _, d := range ext {
			if _ext == strings.ToLower(d) {
				return true
			}
		}
		return false
	}
}

// getFileAll 遍历指定文件夹的目录结构
// 参数path：要遍历的文件夹，注意，末尾不要加斜杠，例如：C:
// 参数files：用于临时存储文件夹或文件信息
// 参数filterFunc：过滤函数，true保留，false不保留
// 返回值1：文件及文件夹的地址信息
// 返回值2：错误
func getFileAllFileFilter(dirPath string, files []string, filterFunc func(filename string) bool) ([]string, error) {
	fInfo, err := os.ReadDir(dirPath)
	if err != nil {
		return files, err
	}
	for _, fi := range fInfo {
		if fi.IsDir() {
			tempDir := filepath.Join(dirPath, fi.Name())
			//files = append(files, tempDir)
			files, _ = getFileAllFileFilter(tempDir, files, filterFunc)
		} else {
			if filterFunc(fi.Name()) {
				files = append(files, filepath.Join(dirPath, fi.Name()))
			}
		}
	}
	return files, nil
}

// SaveAs 文件另存为
// srcFilepath 源文件地址
// newFilepath 另存为的新文件地址
// isDeleteSrc 是否删除掉源文件
func SaveAs(srcFilepath string, newFilepath string, isDeleteSrc bool) error {
	srcFile, err := os.Open(srcFilepath)
	if err != nil {
		return err
	}
	nFile, err := os.Create(newFilepath)
	if err != nil {
		return err
	}
	defer nFile.Close()
	_, err = io.Copy(nFile, srcFile)
	srcFile.Close()
	if isDeleteSrc {
		_ = os.Remove(srcFilepath)
	}
	return err
}

// ReadFileAllRow 读取文件的所有行内容
// srcFilepath 要读取的文件地址
func ReadFileAllRow(srcFilepath string) ([]string, error) {
	file, err := os.Open(srcFilepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		result = append(result, string(line))
	}
	return result, nil
}

// ReadFileAllRowByMinLength 读取文件的所有行内容
// srcFilepath 要读取的文件地址
// minLength 行内容最小长度
func ReadFileAllRowByMinLength(srcFilepath string, minLength int) ([]string, error) {
	file, err := os.Open(srcFilepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		if len(line) < minLength {
			continue
		}
		result = append(result, string(line))
	}
	return result, nil
}

// ReadFileAllRowByMaxLength 读取文件的所有行内容
// srcFilepath 要读取的文件地址
// maxLength 行内容最大长度
func ReadFileAllRowByMaxLength(srcFilepath string, maxLength int) ([]string, error) {
	file, err := os.Open(srcFilepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		if len(line) > maxLength {
			continue
		}
		result = append(result, string(line))
	}
	return result, nil
}

// ReadFileAllRowByLength 读取文件的所有行内容
// srcFilepath 要读取的文件地址
// minLength 行内容最小长度
// maxLength 行内容最大长度
func ReadFileAllRowByLength(srcFilepath string, minLength int, maxLength int) ([]string, error) {
	file, err := os.Open(srcFilepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string
	reader := bufio.NewReader(file)
	length := 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		length = len(line)
		if length > maxLength || length < minLength {
			continue
		}
		result = append(result, string(line))
	}
	return result, nil
}

// FileName 将文件名转为合法的文件名
func FileName(filename string) string {
	filename = strings.ReplaceAll(filename, "/", "_")
	filename = strings.ReplaceAll(filename, `\`, "_")
	filename = strings.ReplaceAll(filename, `:`, "_")
	filename = strings.ReplaceAll(filename, `*`, "_")
	filename = strings.ReplaceAll(filename, `?`, "_")
	filename = strings.ReplaceAll(filename, `"`, "_")
	filename = strings.ReplaceAll(filename, `<`, "_")
	filename = strings.ReplaceAll(filename, `>`, "_")
	filename = strings.ReplaceAll(filename, `|`, "_")
	return filename
}
