/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:23 2019-10-09
 */
package simple_aws_tools

type Options struct {
	AccessKey string
	SecretKey string
	EndPoint string
	Region string
}

type Option func(opts *Options)

func WithAccessKey(key string) Option {
	return func(opts *Options) {
		opts.AccessKey = key
	}
}

func WithSecretKey(key string) Option {
	return func(opts *Options) {
		opts.SecretKey = key
	}
}
func WithEndPoint(endPoint string) Option {
	return func(opts *Options) {
		opts.EndPoint = endPoint
	}
}
func WithRegion(reg string) Option {
	return func(opts *Options) {
		opts.Region = reg
	}
}