// This file is generated by generate-std.joke script. Do not edit manually!

package os

import (
	. "github.com/candid82/joker/core"
	"io/ioutil"
	"os"
)

var __args__P ProcFn = __args_
var args_ Proc = Proc{Fn: __args__P, Name: "args_", Package: "std/os"}

func __args_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 0:
		_res := commandArgs()
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __chdir__P ProcFn = __chdir_
var chdir_ Proc = Proc{Fn: __chdir__P, Name: "chdir_", Package: "std/os"}

func __chdir_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		dirname := ExtractString(_args, 0)
		err := os.Chdir(dirname)
		PanicOnErr(err)
		_res := NIL
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __chmod__P ProcFn = __chmod_
var chmod_ Proc = Proc{Fn: __chmod__P, Name: "chmod_", Package: "std/os"}

func __chmod_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		name := ExtractString(_args, 0)
		mode := ExtractInt(_args, 1)
		err := os.Chmod(name, os.FileMode(mode))
		PanicOnErr(err)
		_res := NIL
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __chown__P ProcFn = __chown_
var chown_ Proc = Proc{Fn: __chown__P, Name: "chown_", Package: "std/os"}

func __chown_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 3:
		name := ExtractString(_args, 0)
		uid := ExtractInt(_args, 1)
		gid := ExtractInt(_args, 2)
		err := os.Chown(name, uid, gid)
		PanicOnErr(err)
		_res := NIL
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __chtimes__P ProcFn = __chtimes_
var chtimes_ Proc = Proc{Fn: __chtimes__P, Name: "chtimes_", Package: "std/os"}

func __chtimes_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 3:
		name := ExtractString(_args, 0)
		atime := ExtractTime(_args, 1)
		mtime := ExtractTime(_args, 2)
		err := os.Chtimes(name, atime, mtime)
		PanicOnErr(err)
		_res := NIL
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __clearenv__P ProcFn = __clearenv_
var clearenv_ Proc = Proc{Fn: __clearenv__P, Name: "clearenv_", Package: "std/os"}

func __clearenv_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 0:
		os.Clearenv()
		_res := NIL
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __close__P ProcFn = __close_
var close_ Proc = Proc{Fn: __close__P, Name: "close_", Package: "std/os"}

func __close_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		f := ExtractFile(_args, 0)
		err := f.Close()
		PanicOnErr(err)
		_res := NIL
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __create__P ProcFn = __create_
var create_ Proc = Proc{Fn: __create__P, Name: "create_", Package: "std/os"}

func __create_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		name := ExtractString(_args, 0)
		_res, err := os.Create(name)
		PanicOnErr(err)
		return MakeFile(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __create_temp__P ProcFn = __create_temp_
var create_temp_ Proc = Proc{Fn: __create_temp__P, Name: "create_temp_", Package: "std/os"}

func __create_temp_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		dir := ExtractString(_args, 0)
		pattern := ExtractString(_args, 1)
		_res, err := ioutil.TempFile(dir, pattern)
		PanicOnErr(err)
		return MakeFile(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __cwd__P ProcFn = __cwd_
var cwd_ Proc = Proc{Fn: __cwd__P, Name: "cwd_", Package: "std/os"}

func __cwd_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 0:
		_res, err := os.Getwd()
		PanicOnErr(err)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __egid__P ProcFn = __egid_
var egid_ Proc = Proc{Fn: __egid__P, Name: "egid_", Package: "std/os"}

func __egid_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 0:
		_res := os.Getegid()
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __env__P ProcFn = __env_
var env_ Proc = Proc{Fn: __env__P, Name: "env_", Package: "std/os"}

func __env_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 0:
		_res := env()
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __euid__P ProcFn = __euid_
var euid_ Proc = Proc{Fn: __euid__P, Name: "euid_", Package: "std/os"}

func __euid_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 0:
		_res := os.Geteuid()
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __exec__P ProcFn = __exec_
var exec_ Proc = Proc{Fn: __exec__P, Name: "exec_", Package: "std/os"}

func __exec_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		name := ExtractString(_args, 0)
		opts := ExtractMap(_args, 1)
		_res := execute(name, opts)
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __executable__P ProcFn = __executable_
var executable_ Proc = Proc{Fn: __executable__P, Name: "executable_", Package: "std/os"}

func __executable_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 0:
		_res, err := os.Executable()
		PanicOnErr(err)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __isexists__P ProcFn = __isexists_
var isexists_ Proc = Proc{Fn: __isexists__P, Name: "isexists_", Package: "std/os"}

func __isexists_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		_res := exists(path)
		return MakeBoolean(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __exit__P ProcFn = __exit_
var exit_ Proc = Proc{Fn: __exit__P, Name: "exit_", Package: "std/os"}

func __exit_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		code := ExtractInt(_args, 0)
		_res := NIL
		ExitJoker(code)
		return _res

	case _c == 0:
		_res := NIL
		ExitJoker(0)
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __expand_env__P ProcFn = __expand_env_
var expand_env_ Proc = Proc{Fn: __expand_env__P, Name: "expand_env_", Package: "std/os"}

func __expand_env_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		s := ExtractString(_args, 0)
		_res := os.ExpandEnv(s)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __get_env__P ProcFn = __get_env_
var get_env_ Proc = Proc{Fn: __get_env__P, Name: "get_env_", Package: "std/os"}

func __get_env_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		key := ExtractString(_args, 0)
		_res := getEnv(key)
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __gid__P ProcFn = __gid_
var gid_ Proc = Proc{Fn: __gid__P, Name: "gid_", Package: "std/os"}

func __gid_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 0:
		_res := os.Getgid()
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __groups__P ProcFn = __groups_
var groups_ Proc = Proc{Fn: __groups__P, Name: "groups_", Package: "std/os"}

func __groups_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 0:
		_res, err := os.Getgroups()
		PanicOnErr(err)
		return MakeIntVector(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __hostname__P ProcFn = __hostname_
var hostname_ Proc = Proc{Fn: __hostname__P, Name: "hostname_", Package: "std/os"}

func __hostname_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 0:
		_res, err := os.Hostname()
		PanicOnErr(err)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __kill_process__P ProcFn = __kill_process_
var kill_process_ Proc = Proc{Fn: __kill_process__P, Name: "kill_process_", Package: "std/os"}

func __kill_process_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		pid := ExtractInt(_args, 0)
		_res := killProcess(pid)
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __lchown__P ProcFn = __lchown_
var lchown_ Proc = Proc{Fn: __lchown__P, Name: "lchown_", Package: "std/os"}

func __lchown_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 3:
		name := ExtractString(_args, 0)
		uid := ExtractInt(_args, 1)
		gid := ExtractInt(_args, 2)
		err := os.Lchown(name, uid, gid)
		PanicOnErr(err)
		_res := NIL
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __link__P ProcFn = __link_
var link_ Proc = Proc{Fn: __link__P, Name: "link_", Package: "std/os"}

func __link_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		oldname := ExtractString(_args, 0)
		newname := ExtractString(_args, 1)
		err := os.Link(oldname, newname)
		PanicOnErr(err)
		_res := NIL
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __ls__P ProcFn = __ls_
var ls_ Proc = Proc{Fn: __ls__P, Name: "ls_", Package: "std/os"}

func __ls_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		dirname := ExtractString(_args, 0)
		_res := readDir(dirname)
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __lstat__P ProcFn = __lstat_
var lstat_ Proc = Proc{Fn: __lstat__P, Name: "lstat_", Package: "std/os"}

func __lstat_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		filename := ExtractString(_args, 0)
		_info, err := os.Lstat(filename)
		PanicOnErr(err)
		_res := FileInfoMap(_info.Name(), _info)
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __mkdir__P ProcFn = __mkdir_
var mkdir_ Proc = Proc{Fn: __mkdir__P, Name: "mkdir_", Package: "std/os"}

func __mkdir_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		name := ExtractString(_args, 0)
		perm := ExtractInt(_args, 1)
		err := os.Mkdir(name, os.FileMode(perm))
		PanicOnErr(err)
		_res := NIL
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __mkdir_all__P ProcFn = __mkdir_all_
var mkdir_all_ Proc = Proc{Fn: __mkdir_all__P, Name: "mkdir_all_", Package: "std/os"}

func __mkdir_all_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		name := ExtractString(_args, 0)
		perm := ExtractInt(_args, 1)
		err := os.MkdirAll(name, os.FileMode(perm))
		PanicOnErr(err)
		_res := NIL
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __mkdir_temp__P ProcFn = __mkdir_temp_
var mkdir_temp_ Proc = Proc{Fn: __mkdir_temp__P, Name: "mkdir_temp_", Package: "std/os"}

func __mkdir_temp_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		dir := ExtractString(_args, 0)
		pattern := ExtractString(_args, 1)
		_res, err := ioutil.TempDir(dir, pattern)
		PanicOnErr(err)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __open__P ProcFn = __open_
var open_ Proc = Proc{Fn: __open__P, Name: "open_", Package: "std/os"}

func __open_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		name := ExtractString(_args, 0)
		_res, err := os.Open(name)
		PanicOnErr(err)
		return MakeFile(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __pagesize__P ProcFn = __pagesize_
var pagesize_ Proc = Proc{Fn: __pagesize__P, Name: "pagesize_", Package: "std/os"}

func __pagesize_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 0:
		_res := os.Getpagesize()
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __ispath_separator__P ProcFn = __ispath_separator_
var ispath_separator_ Proc = Proc{Fn: __ispath_separator__P, Name: "ispath_separator_", Package: "std/os"}

func __ispath_separator_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		c := ExtractChar(_args, 0)
		_res := os.IsPathSeparator(uint8(c))
		return MakeBoolean(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __pid__P ProcFn = __pid_
var pid_ Proc = Proc{Fn: __pid__P, Name: "pid_", Package: "std/os"}

func __pid_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 0:
		_res := os.Getpid()
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __ppid__P ProcFn = __ppid_
var ppid_ Proc = Proc{Fn: __ppid__P, Name: "ppid_", Package: "std/os"}

func __ppid_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 0:
		_res := os.Getppid()
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __read_link__P ProcFn = __read_link_
var read_link_ Proc = Proc{Fn: __read_link__P, Name: "read_link_", Package: "std/os"}

func __read_link_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		name := ExtractString(_args, 0)
		_res, err := os.Readlink(name)
		PanicOnErr(err)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __remove__P ProcFn = __remove_
var remove_ Proc = Proc{Fn: __remove__P, Name: "remove_", Package: "std/os"}

func __remove_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		name := ExtractString(_args, 0)
		err := os.Remove(name)
		PanicOnErr(err)
		_res := NIL
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __remove_all__P ProcFn = __remove_all_
var remove_all_ Proc = Proc{Fn: __remove_all__P, Name: "remove_all_", Package: "std/os"}

func __remove_all_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		path := ExtractString(_args, 0)
		err := os.RemoveAll(path)
		PanicOnErr(err)
		_res := NIL
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __rename__P ProcFn = __rename_
var rename_ Proc = Proc{Fn: __rename__P, Name: "rename_", Package: "std/os"}

func __rename_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		oldpath := ExtractString(_args, 0)
		newpath := ExtractString(_args, 1)
		err := os.Rename(oldpath, newpath)
		PanicOnErr(err)
		_res := NIL
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __set_env__P ProcFn = __set_env_
var set_env_ Proc = Proc{Fn: __set_env__P, Name: "set_env_", Package: "std/os"}

func __set_env_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		key := ExtractString(_args, 0)
		value := ExtractString(_args, 1)
		err := os.Setenv(key, value)
		PanicOnErr(err)
		_res := NIL
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __sh__P ProcFn = __sh_
var sh_ Proc = Proc{Fn: __sh__P, Name: "sh_", Package: "std/os"}

func __sh_(_args []Object) Object {
	_c := len(_args)
	switch {
	case true:
		CheckArity(_args, 1, 999)
		name := ExtractString(_args, 0)
		arguments := ExtractStrings(_args, 1)
		_res := sh("", nil, nil, nil, name, arguments)
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __sh_from__P ProcFn = __sh_from_
var sh_from_ Proc = Proc{Fn: __sh_from__P, Name: "sh_from_", Package: "std/os"}

func __sh_from_(_args []Object) Object {
	_c := len(_args)
	switch {
	case true:
		CheckArity(_args, 2, 999)
		dir := ExtractString(_args, 0)
		name := ExtractString(_args, 1)
		arguments := ExtractStrings(_args, 2)
		_res := sh(dir, nil, nil, nil, name, arguments)
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __signal__P ProcFn = __signal_
var signal_ Proc = Proc{Fn: __signal__P, Name: "signal_", Package: "std/os"}

func __signal_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		pid := ExtractInt(_args, 0)
		signal := ExtractInt(_args, 1)
		_res := sendSignal(pid, signal)
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __start_process__P ProcFn = __start_process_
var start_process_ Proc = Proc{Fn: __start_process__P, Name: "start_process_", Package: "std/os"}

func __start_process_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		name := ExtractString(_args, 0)
		opts := ExtractMap(_args, 1)
		_res := startProcess(name, opts)
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __stat__P ProcFn = __stat_
var stat_ Proc = Proc{Fn: __stat__P, Name: "stat_", Package: "std/os"}

func __stat_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		filename := ExtractString(_args, 0)
		_info, err := os.Stat(filename)
		PanicOnErr(err)
		_res := FileInfoMap(_info.Name(), _info)
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __symlink__P ProcFn = __symlink_
var symlink_ Proc = Proc{Fn: __symlink__P, Name: "symlink_", Package: "std/os"}

func __symlink_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		oldname := ExtractString(_args, 0)
		newname := ExtractString(_args, 1)
		err := os.Symlink(oldname, newname)
		PanicOnErr(err)
		_res := NIL
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __temp_dir__P ProcFn = __temp_dir_
var temp_dir_ Proc = Proc{Fn: __temp_dir__P, Name: "temp_dir_", Package: "std/os"}

func __temp_dir_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 0:
		_res := os.TempDir()
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __truncate__P ProcFn = __truncate_
var truncate_ Proc = Proc{Fn: __truncate__P, Name: "truncate_", Package: "std/os"}

func __truncate_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 2:
		name := ExtractString(_args, 0)
		size := ExtractInt(_args, 1)
		err := os.Truncate(name, int64(size))
		PanicOnErr(err)
		_res := NIL
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __uid__P ProcFn = __uid_
var uid_ Proc = Proc{Fn: __uid__P, Name: "uid_", Package: "std/os"}

func __uid_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 0:
		_res := os.Getuid()
		return MakeInt(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __unset_env__P ProcFn = __unset_env_
var unset_env_ Proc = Proc{Fn: __unset_env__P, Name: "unset_env_", Package: "std/os"}

func __unset_env_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 1:
		key := ExtractString(_args, 0)
		err := os.Unsetenv(key)
		PanicOnErr(err)
		_res := NIL
		return _res

	default:
		PanicArity(_c)
	}
	return NIL
}

var __user_cache_dir__P ProcFn = __user_cache_dir_
var user_cache_dir_ Proc = Proc{Fn: __user_cache_dir__P, Name: "user_cache_dir_", Package: "std/os"}

func __user_cache_dir_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 0:
		_res, err := os.UserCacheDir()
		PanicOnErr(err)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __user_config_dir__P ProcFn = __user_config_dir_
var user_config_dir_ Proc = Proc{Fn: __user_config_dir__P, Name: "user_config_dir_", Package: "std/os"}

func __user_config_dir_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 0:
		_res, err := os.UserConfigDir()
		PanicOnErr(err)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

var __user_home_dir__P ProcFn = __user_home_dir_
var user_home_dir_ Proc = Proc{Fn: __user_home_dir__P, Name: "user_home_dir_", Package: "std/os"}

func __user_home_dir_(_args []Object) Object {
	_c := len(_args)
	switch {
	case _c == 0:
		_res, err := os.UserHomeDir()
		PanicOnErr(err)
		return MakeString(_res)

	default:
		PanicArity(_c)
	}
	return NIL
}

func Init() {

	InternsOrThunks()
}

var osNamespace = GLOBAL_ENV.EnsureSymbolIsLib(MakeSymbol("joker.os"))

func init() {
	osNamespace.Lazy = Init
}
