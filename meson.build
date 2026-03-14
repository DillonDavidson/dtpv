project('dtpv', 'cpp',
	version: '0.0',
	default_options: [
		'cpp_std=c++26',
		'warning_level=3',
	]
)

cmd = run_command('./scripts/grabber.sh', check: true)
srcs = cmd.stdout().strip().split('\n')

deps = []

executable('dtpv', srcs, dependencies: deps)
