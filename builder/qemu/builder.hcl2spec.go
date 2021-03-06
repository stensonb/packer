// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package qemu

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType         *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug               *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce               *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError             *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars            map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars       []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	HTTPDir                   *string           `mapstructure:"http_directory" cty:"http_directory"`
	HTTPPortMin               *int              `mapstructure:"http_port_min" cty:"http_port_min"`
	HTTPPortMax               *int              `mapstructure:"http_port_max" cty:"http_port_max"`
	ISOChecksum               *string           `mapstructure:"iso_checksum" required:"true" cty:"iso_checksum"`
	ISOChecksumURL            *string           `mapstructure:"iso_checksum_url" cty:"iso_checksum_url"`
	ISOChecksumType           *string           `mapstructure:"iso_checksum_type" cty:"iso_checksum_type"`
	RawSingleISOUrl           *string           `mapstructure:"iso_url" required:"true" cty:"iso_url"`
	ISOUrls                   []string          `mapstructure:"iso_urls" cty:"iso_urls"`
	TargetPath                *string           `mapstructure:"iso_target_path" cty:"iso_target_path"`
	TargetExtension           *string           `mapstructure:"iso_target_extension" cty:"iso_target_extension"`
	BootGroupInterval         *string           `mapstructure:"boot_keygroup_interval" cty:"boot_keygroup_interval"`
	BootWait                  *string           `mapstructure:"boot_wait" cty:"boot_wait"`
	BootCommand               []string          `mapstructure:"boot_command" cty:"boot_command"`
	DisableVNC                *bool             `mapstructure:"disable_vnc" cty:"disable_vnc"`
	BootKeyInterval           *string           `mapstructure:"boot_key_interval" cty:"boot_key_interval"`
	ShutdownCommand           *string           `mapstructure:"shutdown_command" required:"false" cty:"shutdown_command"`
	ShutdownTimeout           *string           `mapstructure:"shutdown_timeout" required:"false" cty:"shutdown_timeout"`
	Type                      *string           `mapstructure:"communicator" cty:"communicator"`
	PauseBeforeConnect        *string           `mapstructure:"pause_before_connecting" cty:"pause_before_connecting"`
	SSHHost                   *string           `mapstructure:"ssh_host" cty:"ssh_host"`
	SSHPort                   *int              `mapstructure:"ssh_port" cty:"ssh_port"`
	SSHUsername               *string           `mapstructure:"ssh_username" cty:"ssh_username"`
	SSHPassword               *string           `mapstructure:"ssh_password" cty:"ssh_password"`
	SSHKeyPairName            *string           `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   *string           `mapstructure:"temporary_key_pair_name" cty:"temporary_key_pair_name"`
	SSHClearAuthorizedKeys    *bool             `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile         *string           `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file"`
	SSHPty                    *bool             `mapstructure:"ssh_pty" cty:"ssh_pty"`
	SSHTimeout                *string           `mapstructure:"ssh_timeout" cty:"ssh_timeout"`
	SSHAgentAuth              *bool             `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth"`
	SSHDisableAgentForwarding *bool             `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      *int              `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts"`
	SSHBastionHost            *string           `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host"`
	SSHBastionPort            *int              `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port"`
	SSHBastionAgentAuth       *bool             `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth"`
	SSHBastionUsername        *string           `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username"`
	SSHBastionPassword        *string           `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password"`
	SSHBastionPrivateKeyFile  *string           `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod     *string           `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method"`
	SSHProxyHost              *string           `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host"`
	SSHProxyPort              *int              `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port"`
	SSHProxyUsername          *string           `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username"`
	SSHProxyPassword          *string           `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password"`
	SSHKeepAliveInterval      *string           `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       *string           `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string          `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string          `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels"`
	SSHPublicKey              []byte            `cty:"ssh_public_key"`
	SSHPrivateKey             []byte            `cty:"ssh_private_key"`
	WinRMUser                 *string           `mapstructure:"winrm_username" cty:"winrm_username"`
	WinRMPassword             *string           `mapstructure:"winrm_password" cty:"winrm_password"`
	WinRMHost                 *string           `mapstructure:"winrm_host" cty:"winrm_host"`
	WinRMPort                 *int              `mapstructure:"winrm_port" cty:"winrm_port"`
	WinRMTimeout              *string           `mapstructure:"winrm_timeout" cty:"winrm_timeout"`
	WinRMUseSSL               *bool             `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl"`
	WinRMInsecure             *bool             `mapstructure:"winrm_insecure" cty:"winrm_insecure"`
	WinRMUseNTLM              *bool             `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm"`
	FloppyFiles               []string          `mapstructure:"floppy_files" cty:"floppy_files"`
	FloppyDirectories         []string          `mapstructure:"floppy_dirs" cty:"floppy_dirs"`
	FloppyLabel               *string           `mapstructure:"floppy_label" cty:"floppy_label"`
	ISOSkipCache              *bool             `mapstructure:"iso_skip_cache" required:"false" cty:"iso_skip_cache"`
	Accelerator               *string           `mapstructure:"accelerator" required:"false" cty:"accelerator"`
	AdditionalDiskSize        []string          `mapstructure:"disk_additional_size" required:"false" cty:"disk_additional_size"`
	CpuCount                  *int              `mapstructure:"cpus" required:"false" cty:"cpus"`
	DiskInterface             *string           `mapstructure:"disk_interface" required:"false" cty:"disk_interface"`
	DiskSize                  *string           `mapstructure:"disk_size" required:"false" cty:"disk_size"`
	DiskCache                 *string           `mapstructure:"disk_cache" required:"false" cty:"disk_cache"`
	DiskDiscard               *string           `mapstructure:"disk_discard" required:"false" cty:"disk_discard"`
	DetectZeroes              *string           `mapstructure:"disk_detect_zeroes" required:"false" cty:"disk_detect_zeroes"`
	SkipCompaction            *bool             `mapstructure:"skip_compaction" required:"false" cty:"skip_compaction"`
	DiskCompression           *bool             `mapstructure:"disk_compression" required:"false" cty:"disk_compression"`
	Format                    *string           `mapstructure:"format" required:"false" cty:"format"`
	Headless                  *bool             `mapstructure:"headless" required:"false" cty:"headless"`
	DiskImage                 *bool             `mapstructure:"disk_image" required:"false" cty:"disk_image"`
	UseBackingFile            *bool             `mapstructure:"use_backing_file" required:"false" cty:"use_backing_file"`
	MachineType               *string           `mapstructure:"machine_type" required:"false" cty:"machine_type"`
	MemorySize                *int              `mapstructure:"memory" required:"false" cty:"memory"`
	NetDevice                 *string           `mapstructure:"net_device" required:"false" cty:"net_device"`
	OutputDir                 *string           `mapstructure:"output_directory" required:"false" cty:"output_directory"`
	QemuArgs                  [][]string        `mapstructure:"qemuargs" required:"false" cty:"qemuargs"`
	QemuBinary                *string           `mapstructure:"qemu_binary" required:"false" cty:"qemu_binary"`
	QMPEnable                 *bool             `mapstructure:"qmp_enable" required:"false" cty:"qmp_enable"`
	QMPSocketPath             *string           `mapstructure:"qmp_socket_path" required:"false" cty:"qmp_socket_path"`
	SSHHostPortMin            *int              `mapstructure:"ssh_host_port_min" required:"false" cty:"ssh_host_port_min"`
	SSHHostPortMax            *int              `mapstructure:"ssh_host_port_max" required:"false" cty:"ssh_host_port_max"`
	UseDefaultDisplay         *bool             `mapstructure:"use_default_display" required:"false" cty:"use_default_display"`
	VNCBindAddress            *string           `mapstructure:"vnc_bind_address" required:"false" cty:"vnc_bind_address"`
	VNCUsePassword            *bool             `mapstructure:"vnc_use_password" required:"false" cty:"vnc_use_password"`
	VNCPortMin                *int              `mapstructure:"vnc_port_min" required:"false" cty:"vnc_port_min"`
	VNCPortMax                *int              `mapstructure:"vnc_port_max" cty:"vnc_port_max"`
	VMName                    *string           `mapstructure:"vm_name" required:"false" cty:"vm_name"`
	SSHWaitTimeout            *string           `mapstructure:"ssh_wait_timeout" required:"false" cty:"ssh_wait_timeout"`
	RunOnce                   *bool             `mapstructure:"run_once" cty:"run_once"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{} { return new(FlatConfig) }

// HCL2Spec returns the hcldec.Spec of a FlatConfig.
// This spec is used by HCL to read the fields of FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":            &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":          &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":                 &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                 &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":              &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":        &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables":   &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"http_directory":               &hcldec.AttrSpec{Name: "http_directory", Type: cty.String, Required: false},
		"http_port_min":                &hcldec.AttrSpec{Name: "http_port_min", Type: cty.Number, Required: false},
		"http_port_max":                &hcldec.AttrSpec{Name: "http_port_max", Type: cty.Number, Required: false},
		"iso_checksum":                 &hcldec.AttrSpec{Name: "iso_checksum", Type: cty.String, Required: false},
		"iso_checksum_url":             &hcldec.AttrSpec{Name: "iso_checksum_url", Type: cty.String, Required: false},
		"iso_checksum_type":            &hcldec.AttrSpec{Name: "iso_checksum_type", Type: cty.String, Required: false},
		"iso_url":                      &hcldec.AttrSpec{Name: "iso_url", Type: cty.String, Required: false},
		"iso_urls":                     &hcldec.AttrSpec{Name: "iso_urls", Type: cty.List(cty.String), Required: false},
		"iso_target_path":              &hcldec.AttrSpec{Name: "iso_target_path", Type: cty.String, Required: false},
		"iso_target_extension":         &hcldec.AttrSpec{Name: "iso_target_extension", Type: cty.String, Required: false},
		"boot_keygroup_interval":       &hcldec.AttrSpec{Name: "boot_keygroup_interval", Type: cty.String, Required: false},
		"boot_wait":                    &hcldec.AttrSpec{Name: "boot_wait", Type: cty.String, Required: false},
		"boot_command":                 &hcldec.AttrSpec{Name: "boot_command", Type: cty.List(cty.String), Required: false},
		"disable_vnc":                  &hcldec.AttrSpec{Name: "disable_vnc", Type: cty.Bool, Required: false},
		"boot_key_interval":            &hcldec.AttrSpec{Name: "boot_key_interval", Type: cty.String, Required: false},
		"shutdown_command":             &hcldec.AttrSpec{Name: "shutdown_command", Type: cty.String, Required: false},
		"shutdown_timeout":             &hcldec.AttrSpec{Name: "shutdown_timeout", Type: cty.String, Required: false},
		"communicator":                 &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":      &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                     &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                     &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                 &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                 &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":             &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":      &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"ssh_clear_authorized_keys":    &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_private_key_file":         &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_pty":                      &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                  &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":               &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding": &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":       &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":             &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":             &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":       &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":         &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":         &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_private_key_file": &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":     &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":               &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":               &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":           &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":           &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":      &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":       &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":           &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":            &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":               &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":              &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":               &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":               &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                   &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_port":                   &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":               &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":               &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"floppy_files":                 &hcldec.AttrSpec{Name: "floppy_files", Type: cty.List(cty.String), Required: false},
		"floppy_dirs":                  &hcldec.AttrSpec{Name: "floppy_dirs", Type: cty.List(cty.String), Required: false},
		"floppy_label":                 &hcldec.AttrSpec{Name: "floppy_label", Type: cty.String, Required: false},
		"iso_skip_cache":               &hcldec.AttrSpec{Name: "iso_skip_cache", Type: cty.Bool, Required: false},
		"accelerator":                  &hcldec.AttrSpec{Name: "accelerator", Type: cty.String, Required: false},
		"disk_additional_size":         &hcldec.AttrSpec{Name: "disk_additional_size", Type: cty.List(cty.String), Required: false},
		"cpus":                         &hcldec.AttrSpec{Name: "cpus", Type: cty.Number, Required: false},
		"disk_interface":               &hcldec.AttrSpec{Name: "disk_interface", Type: cty.String, Required: false},
		"disk_size":                    &hcldec.AttrSpec{Name: "disk_size", Type: cty.String, Required: false},
		"disk_cache":                   &hcldec.AttrSpec{Name: "disk_cache", Type: cty.String, Required: false},
		"disk_discard":                 &hcldec.AttrSpec{Name: "disk_discard", Type: cty.String, Required: false},
		"disk_detect_zeroes":           &hcldec.AttrSpec{Name: "disk_detect_zeroes", Type: cty.String, Required: false},
		"skip_compaction":              &hcldec.AttrSpec{Name: "skip_compaction", Type: cty.Bool, Required: false},
		"disk_compression":             &hcldec.AttrSpec{Name: "disk_compression", Type: cty.Bool, Required: false},
		"format":                       &hcldec.AttrSpec{Name: "format", Type: cty.String, Required: false},
		"headless":                     &hcldec.AttrSpec{Name: "headless", Type: cty.Bool, Required: false},
		"disk_image":                   &hcldec.AttrSpec{Name: "disk_image", Type: cty.Bool, Required: false},
		"use_backing_file":             &hcldec.AttrSpec{Name: "use_backing_file", Type: cty.Bool, Required: false},
		"machine_type":                 &hcldec.AttrSpec{Name: "machine_type", Type: cty.String, Required: false},
		"memory":                       &hcldec.AttrSpec{Name: "memory", Type: cty.Number, Required: false},
		"net_device":                   &hcldec.AttrSpec{Name: "net_device", Type: cty.String, Required: false},
		"output_directory":             &hcldec.AttrSpec{Name: "output_directory", Type: cty.String, Required: false},
		"qemuargs":                     &hcldec.BlockListSpec{TypeName: "qemuargs", Nested: &hcldec.AttrSpec{Name: "qemuargs", Type: cty.List(cty.String), Required: false}},
		"qemu_binary":                  &hcldec.AttrSpec{Name: "qemu_binary", Type: cty.String, Required: false},
		"qmp_enable":                   &hcldec.AttrSpec{Name: "qmp_enable", Type: cty.Bool, Required: false},
		"qmp_socket_path":              &hcldec.AttrSpec{Name: "qmp_socket_path", Type: cty.String, Required: false},
		"ssh_host_port_min":            &hcldec.AttrSpec{Name: "ssh_host_port_min", Type: cty.Number, Required: false},
		"ssh_host_port_max":            &hcldec.AttrSpec{Name: "ssh_host_port_max", Type: cty.Number, Required: false},
		"use_default_display":          &hcldec.AttrSpec{Name: "use_default_display", Type: cty.Bool, Required: false},
		"vnc_bind_address":             &hcldec.AttrSpec{Name: "vnc_bind_address", Type: cty.String, Required: false},
		"vnc_use_password":             &hcldec.AttrSpec{Name: "vnc_use_password", Type: cty.Bool, Required: false},
		"vnc_port_min":                 &hcldec.AttrSpec{Name: "vnc_port_min", Type: cty.Number, Required: false},
		"vnc_port_max":                 &hcldec.AttrSpec{Name: "vnc_port_max", Type: cty.Number, Required: false},
		"vm_name":                      &hcldec.AttrSpec{Name: "vm_name", Type: cty.String, Required: false},
		"ssh_wait_timeout":             &hcldec.AttrSpec{Name: "ssh_wait_timeout", Type: cty.String, Required: false},
		"run_once":                     &hcldec.AttrSpec{Name: "run_once", Type: cty.Bool, Required: false},
	}
	return s
}
