project('dtpv', 'cpp',
	version: '0.0',
	default_options: [
		'cpp_std=c++26',
		'warning_level=3',
		'optimization=3',
	]
)

srcs = [
	'src/main.cpp',
	'src/file_type.cpp',
	'src/video.cpp',
	'src/text.cpp',
	'src/markdown.cpp',
	'src/image.cpp',
	'src/pdf.cpp',
	'src/cache.cpp',
]

deps = []

executable('dtpv', srcs, dependencies: deps)
