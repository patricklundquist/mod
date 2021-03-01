// Code generated by protoc-gen-cobra. DO NOT EDIT.

package v2

import (
	cobra "github.com/spf13/cobra"
	client "go.amplifyedge.org/protoc-gen-cobra/client"
	flag "go.amplifyedge.org/protoc-gen-cobra/flag"
	iocodec "go.amplifyedge.org/protoc-gen-cobra/iocodec"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func SurveyServiceClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("SurveyService"),
		Short:  "SurveyService service client",
		Long:   "",
		Hidden: false,
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_SurveyServiceNewSurveyProjectCommand(cfg),
		_SurveyServiceGetSurveyProjectCommand(cfg),
		_SurveyServiceListSurveyProjectCommand(cfg),
		_SurveyServiceUpdateSurveyProjectCommand(cfg),
		_SurveyServiceDeleteSurveyProjectCommand(cfg),
		_SurveyServiceNewSurveyUserCommand(cfg),
		_SurveyServiceGetSurveyUserCommand(cfg),
		_SurveyServiceListSurveyUserCommand(cfg),
		_SurveyServiceUpdateSurveyUserCommand(cfg),
		_SurveyServiceDeleteSurveyUserCommand(cfg),
		_SurveyServiceGetProjectStatisticsCommand(cfg),
		_SurveyServiceNewDiscoProjectCommand(cfg),
		_SurveyServiceGetDiscoProjectCommand(cfg),
		_SurveyServiceListDiscoProjectCommand(cfg),
		_SurveyServiceUpdateDiscoProjectCommand(cfg),
		_SurveyServiceDeleteDiscoProjectCommand(cfg),
		_SurveyServiceGenTempIdCommand(cfg),
	)
	return cmd
}

func _SurveyServiceNewSurveyProjectCommand(cfg *client.Config) *cobra.Command {
	req := &NewSurveyProjectRequest{}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("NewSurveyProject"),
		Short:  "NewSurveyProject RPC client",
		Long:   "Projects",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "NewSurveyProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &NewSurveyProjectRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.NewSurveyProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SysAccountProjectRefId, cfg.FlagNamer("SysAccountProjectRefId"), "", "@inject_tag: fake:\"skip\" yaml:\"sys_account_project_ref_id,omitempty\"")
	cmd.PersistentFlags().StringVar(&req.SysAccountProjectRefName, cfg.FlagNamer("SysAccountProjectRefName"), "", "@inject_tag: fake:\"{nameseq:sys_account_project,true,sys_account_project,true,false}\" yaml:\"sys_account_project_ref_name\"")
	cmd.PersistentFlags().StringVar(&req.SurveyProjectName, cfg.FlagNamer("SurveyProjectName"), "", "@inject_tag: fake:\"{nameseq:survey_project,false,none,false,false}\" yaml:\"survey_project_name\"")

	return cmd
}

func _SurveyServiceGetSurveyProjectCommand(cfg *client.Config) *cobra.Command {
	req := &IdRequest{}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("GetSurveyProject"),
		Short:  "GetSurveyProject RPC client",
		Long:   "",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "GetSurveyProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &IdRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetSurveyProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SurveyUserId, cfg.FlagNamer("SurveyUserId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountProjectId, cfg.FlagNamer("SysAccountProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SurveyProjectId, cfg.FlagNamer("SurveyProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountAccountId, cfg.FlagNamer("SysAccountAccountId"), "", "")
	cmd.PersistentFlags().StringVar(&req.DiscoProjectId, cfg.FlagNamer("DiscoProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountOrgId, cfg.FlagNamer("SysAccountOrgId"), "", "")

	return cmd
}

func _SurveyServiceListSurveyProjectCommand(cfg *client.Config) *cobra.Command {
	req := &ListRequest{
		IdRequest: &IdRequest{},
	}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("ListSurveyProject"),
		Short:  "ListSurveyProject RPC client",
		Long:   "",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "ListSurveyProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &ListRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.ListSurveyProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.IdRequest.SurveyUserId, cfg.FlagNamer("IdRequest SurveyUserId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SysAccountProjectId, cfg.FlagNamer("IdRequest SysAccountProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SurveyProjectId, cfg.FlagNamer("IdRequest SurveyProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SysAccountAccountId, cfg.FlagNamer("IdRequest SysAccountAccountId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.DiscoProjectId, cfg.FlagNamer("IdRequest DiscoProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SysAccountOrgId, cfg.FlagNamer("IdRequest SysAccountOrgId"), "", "")
	cmd.PersistentFlags().Int64Var(&req.PerPageEntries, cfg.FlagNamer("PerPageEntries"), 0, "")
	cmd.PersistentFlags().StringVar(&req.OrderBy, cfg.FlagNamer("OrderBy"), "", "")
	cmd.PersistentFlags().StringVar(&req.CurrentPageId, cfg.FlagNamer("CurrentPageId"), "", "")
	flag.BytesBase64Var(cmd.PersistentFlags(), &req.Filters, cfg.FlagNamer("Filters"), "")
	cmd.PersistentFlags().BoolVar(&req.IsDescending, cfg.FlagNamer("IsDescending"), false, "")
	cmd.PersistentFlags().StringVar(&req.Matcher, cfg.FlagNamer("Matcher"), "", "")

	return cmd
}

func _SurveyServiceUpdateSurveyProjectCommand(cfg *client.Config) *cobra.Command {
	req := &UpdateSurveyProjectRequest{}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("UpdateSurveyProject"),
		Short:  "UpdateSurveyProject RPC client",
		Long:   "",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "UpdateSurveyProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &UpdateSurveyProjectRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.UpdateSurveyProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SurveyProjectId, cfg.FlagNamer("SurveyProjectId"), "", "")

	return cmd
}

func _SurveyServiceDeleteSurveyProjectCommand(cfg *client.Config) *cobra.Command {
	req := &IdRequest{}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("DeleteSurveyProject"),
		Short:  "DeleteSurveyProject RPC client",
		Long:   "",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "DeleteSurveyProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &IdRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.DeleteSurveyProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SurveyUserId, cfg.FlagNamer("SurveyUserId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountProjectId, cfg.FlagNamer("SysAccountProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SurveyProjectId, cfg.FlagNamer("SurveyProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountAccountId, cfg.FlagNamer("SysAccountAccountId"), "", "")
	cmd.PersistentFlags().StringVar(&req.DiscoProjectId, cfg.FlagNamer("DiscoProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountOrgId, cfg.FlagNamer("SysAccountOrgId"), "", "")

	return cmd
}

func _SurveyServiceNewSurveyUserCommand(cfg *client.Config) *cobra.Command {
	req := &NewSurveyUserRequest{}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("NewSurveyUser"),
		Short:  "NewSurveyUser RPC client",
		Long:   "Users",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "NewSurveyUser"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &NewSurveyUserRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.NewSurveyUser(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SurveyProjectRefId, cfg.FlagNamer("SurveyProjectRefId"), "", "@inject_tag: fake:\"skip\" yaml:\"survey_project_ref_id,omitempty\"")
	cmd.PersistentFlags().StringVar(&req.SysAccountUserRefId, cfg.FlagNamer("SysAccountUserRefId"), "", "@inject_tag: fake:\"skip\" yaml:\"sys_account_user_ref_id,omitempty\"")
	cmd.PersistentFlags().StringVar(&req.SurveyProjectRefName, cfg.FlagNamer("SurveyProjectRefName"), "", "@inject_tag: fake:\"{nameseq:survey_project,true,survey_project,false,false}\" yaml:\"survey_project_ref_name\"")
	cmd.PersistentFlags().StringVar(&req.SysAccountUserRefName, cfg.FlagNamer("SysAccountUserRefName"), "", "@inject_tag: fake:\"{mailseq:user,true,sys_account_email,true}\" yaml:\"sys_account_user_ref_name\"")
	cmd.PersistentFlags().StringVar(&req.SurveyUserName, cfg.FlagNamer("SurveyUserName"), "", "@inject_tag: fake:\"{nameseq:survey_user,false,none,false,false}\" yaml:\"survey_user_name\"")

	return cmd
}

func _SurveyServiceGetSurveyUserCommand(cfg *client.Config) *cobra.Command {
	req := &IdRequest{}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("GetSurveyUser"),
		Short:  "GetSurveyUser RPC client",
		Long:   "",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "GetSurveyUser"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &IdRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetSurveyUser(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SurveyUserId, cfg.FlagNamer("SurveyUserId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountProjectId, cfg.FlagNamer("SysAccountProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SurveyProjectId, cfg.FlagNamer("SurveyProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountAccountId, cfg.FlagNamer("SysAccountAccountId"), "", "")
	cmd.PersistentFlags().StringVar(&req.DiscoProjectId, cfg.FlagNamer("DiscoProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountOrgId, cfg.FlagNamer("SysAccountOrgId"), "", "")

	return cmd
}

func _SurveyServiceListSurveyUserCommand(cfg *client.Config) *cobra.Command {
	req := &ListRequest{
		IdRequest: &IdRequest{},
	}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("ListSurveyUser"),
		Short:  "ListSurveyUser RPC client",
		Long:   "",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "ListSurveyUser"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &ListRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.ListSurveyUser(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.IdRequest.SurveyUserId, cfg.FlagNamer("IdRequest SurveyUserId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SysAccountProjectId, cfg.FlagNamer("IdRequest SysAccountProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SurveyProjectId, cfg.FlagNamer("IdRequest SurveyProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SysAccountAccountId, cfg.FlagNamer("IdRequest SysAccountAccountId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.DiscoProjectId, cfg.FlagNamer("IdRequest DiscoProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SysAccountOrgId, cfg.FlagNamer("IdRequest SysAccountOrgId"), "", "")
	cmd.PersistentFlags().Int64Var(&req.PerPageEntries, cfg.FlagNamer("PerPageEntries"), 0, "")
	cmd.PersistentFlags().StringVar(&req.OrderBy, cfg.FlagNamer("OrderBy"), "", "")
	cmd.PersistentFlags().StringVar(&req.CurrentPageId, cfg.FlagNamer("CurrentPageId"), "", "")
	flag.BytesBase64Var(cmd.PersistentFlags(), &req.Filters, cfg.FlagNamer("Filters"), "")
	cmd.PersistentFlags().BoolVar(&req.IsDescending, cfg.FlagNamer("IsDescending"), false, "")
	cmd.PersistentFlags().StringVar(&req.Matcher, cfg.FlagNamer("Matcher"), "", "")

	return cmd
}

func _SurveyServiceUpdateSurveyUserCommand(cfg *client.Config) *cobra.Command {
	req := &UpdateSurveyUserRequest{}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("UpdateSurveyUser"),
		Short:  "UpdateSurveyUser RPC client",
		Long:   "",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "UpdateSurveyUser"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &UpdateSurveyUserRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.UpdateSurveyUser(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SurveyUserId, cfg.FlagNamer("SurveyUserId"), "", "")

	return cmd
}

func _SurveyServiceDeleteSurveyUserCommand(cfg *client.Config) *cobra.Command {
	req := &IdRequest{}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("DeleteSurveyUser"),
		Short:  "DeleteSurveyUser RPC client",
		Long:   "",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "DeleteSurveyUser"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &IdRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.DeleteSurveyUser(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SurveyUserId, cfg.FlagNamer("SurveyUserId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountProjectId, cfg.FlagNamer("SysAccountProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SurveyProjectId, cfg.FlagNamer("SurveyProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountAccountId, cfg.FlagNamer("SysAccountAccountId"), "", "")
	cmd.PersistentFlags().StringVar(&req.DiscoProjectId, cfg.FlagNamer("DiscoProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountOrgId, cfg.FlagNamer("SysAccountOrgId"), "", "")

	return cmd
}

func _SurveyServiceGetProjectStatisticsCommand(cfg *client.Config) *cobra.Command {
	req := &StatisticRequest{}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("GetProjectStatistics"),
		Short:  "GetProjectStatistics RPC client",
		Long:   "",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "GetProjectStatistics"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &StatisticRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetProjectStatistics(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.TableName, cfg.FlagNamer("TableName"), "", "")
	cmd.PersistentFlags().BoolVar(&req.IncludeCounts, cfg.FlagNamer("IncludeCounts"), false, "")
	flag.BytesBase64Var(cmd.PersistentFlags(), &req.Filters, cfg.FlagNamer("Filters"), "")
	cmd.PersistentFlags().StringVar(&req.CurrentPageId, cfg.FlagNamer("CurrentPageId"), "", "")
	cmd.PersistentFlags().Int64Var(&req.Limit, cfg.FlagNamer("Limit"), 0, "")
	cmd.PersistentFlags().StringVar(&req.OrderBy, cfg.FlagNamer("OrderBy"), "", "")
	cmd.PersistentFlags().BoolVar(&req.IsDescending, cfg.FlagNamer("IsDescending"), false, "")
	cmd.PersistentFlags().Int64Var(&req.PerPageEntries, cfg.FlagNamer("PerPageEntries"), 0, "")

	return cmd
}

func _SurveyServiceNewDiscoProjectCommand(cfg *client.Config) *cobra.Command {
	req := &NewDiscoProjectRequest{}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("NewDiscoProject"),
		Short:  "NewDiscoProject RPC client",
		Long:   "DiscoProjects",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "NewDiscoProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &NewDiscoProjectRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.NewDiscoProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SysAccountProjectRefId, cfg.FlagNamer("SysAccountProjectRefId"), "", "@inject_tag: fake:\"skip\" yaml:\"sys_account_project_ref_id,omitempty\"")
	cmd.PersistentFlags().StringVar(&req.SysAccountOrgRefId, cfg.FlagNamer("SysAccountOrgRefId"), "", "@inject_tag: fake:\"skip\" yaml:\"sys_account_org_ref_id,omitempty\"")
	cmd.PersistentFlags().StringVar(&req.Goal, cfg.FlagNamer("Goal"), "", "@inject_tag: fake:\"{sentence:8}\" yaml:\"goal,omitempty\"")
	cmd.PersistentFlags().Uint64Var(&req.AlreadyPledged, cfg.FlagNamer("AlreadyPledged"), 0, "@inject_tag: fake:\"skip\" yaml:\"already_pledged,omitempty\"")
	cmd.PersistentFlags().Int64Var(&req.ActionTimeNano, cfg.FlagNamer("ActionTimeNano"), 0, "@inject_tag: fake:\"{randomts}\" yaml:\"action_time\"")
	cmd.PersistentFlags().StringVar(&req.ActionLocation, cfg.FlagNamer("ActionLocation"), "", "@inject_tag: fake:\"{city}\" yaml:\"action_location\"")
	cmd.PersistentFlags().Uint64Var(&req.MinPioneers, cfg.FlagNamer("MinPioneers"), 0, "@inject_tag: fake:\"{number:100,1000}\" yaml:\"min_pioneers\"")
	cmd.PersistentFlags().Uint64Var(&req.MinRebelsMedia, cfg.FlagNamer("MinRebelsMedia"), 0, "@inject_tag: fake:\"{number:100,150}\" yaml:\"min_rebels_media\"")
	cmd.PersistentFlags().Uint64Var(&req.MinRebelsToWin, cfg.FlagNamer("MinRebelsToWin"), 0, "@inject_tag: fake:\"{number:100,150}\" yaml:\"min_rebels_to_win\"")
	cmd.PersistentFlags().StringVar(&req.ActionLength, cfg.FlagNamer("ActionLength"), "", "@inject_tag: fake:\"{randomstring:[14 days, 13 weeks, 12 months]}\" yaml:\"action_length\"")
	cmd.PersistentFlags().StringVar(&req.ActionType, cfg.FlagNamer("ActionType"), "", "@inject_tag: fake:\"{randomstring:[environment,global_campaign,poverty]}\" yaml:\"action_type\"")
	cmd.PersistentFlags().StringVar(&req.Category, cfg.FlagNamer("Category"), "", "@inject_tag: fake:\"{randomstring:[environment,global_campaign,poverty]}\" yaml:\"category\"")
	cmd.PersistentFlags().StringVar(&req.Contact, cfg.FlagNamer("Contact"), "", "@inject_tag: fake:\"{email}\" yaml:\"contact\"")
	cmd.PersistentFlags().StringVar(&req.HistPrecedents, cfg.FlagNamer("HistPrecedents"), "", "@inject_tag: fake:\"{sentence:8}\" yaml:\"hist_precedents\"")
	cmd.PersistentFlags().StringVar(&req.Strategy, cfg.FlagNamer("Strategy"), "", "@inject_tag: fake:\"{sentence:8}\" yaml:\"strategy\"")
	cmd.PersistentFlags().StringSliceVar(&req.VideoUrl, cfg.FlagNamer("VideoUrl"), nil, "@inject_tag: fake:\"{randomyt}\" fakesize:\"2\" yaml:\"video_url\"")
	cmd.PersistentFlags().StringVar(&req.UnitOfMeasures, cfg.FlagNamer("UnitOfMeasures"), "", "@inject_tag: fake:\"{randomstring:[days,weeks,months]}\" yaml:\"unit_of_measures\"")
	cmd.PersistentFlags().StringVar(&req.SysAccountProjectRefName, cfg.FlagNamer("SysAccountProjectRefName"), "", "@inject_tag: fake:\"{nameseq:sys_account_project,true,sys_account_project,true,false}\" yaml:\"sys_account_project_ref_name\"")
	cmd.PersistentFlags().StringVar(&req.SysAccountOrgRefName, cfg.FlagNamer("SysAccountOrgRefName"), "", "@inject_tag: fake:\"{nameseq:sys_account_org,true,sys_account_org,false,false}\" yaml:\"sys_account_org_ref_name,omitempty\"")
	cmd.PersistentFlags().StringSliceVar(&req.ImageFilepath, cfg.FlagNamer("ImageFilepath"), nil, "@inject_tag: fake:\"skip\" yaml:\"image_filepath,omitempty\"")
	cmd.PersistentFlags().StringSliceVar(&req.ImageUploadArrays, cfg.FlagNamer("ImageUploadArrays"), nil, "@inject_tag: fake:\"{logogen:300}\" yaml:\"image_upload_arrays,omitempty\"")

	return cmd
}

func _SurveyServiceGetDiscoProjectCommand(cfg *client.Config) *cobra.Command {
	req := &IdRequest{}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("GetDiscoProject"),
		Short:  "GetDiscoProject RPC client",
		Long:   "",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "GetDiscoProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &IdRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetDiscoProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SurveyUserId, cfg.FlagNamer("SurveyUserId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountProjectId, cfg.FlagNamer("SysAccountProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SurveyProjectId, cfg.FlagNamer("SurveyProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountAccountId, cfg.FlagNamer("SysAccountAccountId"), "", "")
	cmd.PersistentFlags().StringVar(&req.DiscoProjectId, cfg.FlagNamer("DiscoProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountOrgId, cfg.FlagNamer("SysAccountOrgId"), "", "")

	return cmd
}

func _SurveyServiceListDiscoProjectCommand(cfg *client.Config) *cobra.Command {
	req := &ListRequest{
		IdRequest: &IdRequest{},
	}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("ListDiscoProject"),
		Short:  "ListDiscoProject RPC client",
		Long:   "",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "ListDiscoProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &ListRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.ListDiscoProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.IdRequest.SurveyUserId, cfg.FlagNamer("IdRequest SurveyUserId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SysAccountProjectId, cfg.FlagNamer("IdRequest SysAccountProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SurveyProjectId, cfg.FlagNamer("IdRequest SurveyProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SysAccountAccountId, cfg.FlagNamer("IdRequest SysAccountAccountId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.DiscoProjectId, cfg.FlagNamer("IdRequest DiscoProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SysAccountOrgId, cfg.FlagNamer("IdRequest SysAccountOrgId"), "", "")
	cmd.PersistentFlags().Int64Var(&req.PerPageEntries, cfg.FlagNamer("PerPageEntries"), 0, "")
	cmd.PersistentFlags().StringVar(&req.OrderBy, cfg.FlagNamer("OrderBy"), "", "")
	cmd.PersistentFlags().StringVar(&req.CurrentPageId, cfg.FlagNamer("CurrentPageId"), "", "")
	flag.BytesBase64Var(cmd.PersistentFlags(), &req.Filters, cfg.FlagNamer("Filters"), "")
	cmd.PersistentFlags().BoolVar(&req.IsDescending, cfg.FlagNamer("IsDescending"), false, "")
	cmd.PersistentFlags().StringVar(&req.Matcher, cfg.FlagNamer("Matcher"), "", "")

	return cmd
}

func _SurveyServiceUpdateDiscoProjectCommand(cfg *client.Config) *cobra.Command {
	req := &UpdateDiscoProjectRequest{}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("UpdateDiscoProject"),
		Short:  "UpdateDiscoProject RPC client",
		Long:   "",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "UpdateDiscoProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &UpdateDiscoProjectRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.UpdateDiscoProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.ProjectId, cfg.FlagNamer("ProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.Goal, cfg.FlagNamer("Goal"), "", "")
	cmd.PersistentFlags().Uint64Var(&req.AlreadyPledged, cfg.FlagNamer("AlreadyPledged"), 0, "")
	flag.TimestampVar(cmd.PersistentFlags(), &req.ActionTime, cfg.FlagNamer("ActionTime"), "")
	cmd.PersistentFlags().StringVar(&req.ActionLocation, cfg.FlagNamer("ActionLocation"), "", "")
	cmd.PersistentFlags().Uint64Var(&req.MinPioneers, cfg.FlagNamer("MinPioneers"), 0, "")
	cmd.PersistentFlags().Uint64Var(&req.MinRebelsMedia, cfg.FlagNamer("MinRebelsMedia"), 0, "")
	cmd.PersistentFlags().Uint64Var(&req.MinRebelsToWin, cfg.FlagNamer("MinRebelsToWin"), 0, "")
	cmd.PersistentFlags().StringVar(&req.ActionLength, cfg.FlagNamer("ActionLength"), "", "")
	cmd.PersistentFlags().StringVar(&req.ActionType, cfg.FlagNamer("ActionType"), "", "")
	cmd.PersistentFlags().StringVar(&req.Category, cfg.FlagNamer("Category"), "", "")
	cmd.PersistentFlags().StringVar(&req.Contact, cfg.FlagNamer("Contact"), "", "")
	cmd.PersistentFlags().StringVar(&req.HistPrecedents, cfg.FlagNamer("HistPrecedents"), "", "")
	cmd.PersistentFlags().StringVar(&req.Strategy, cfg.FlagNamer("Strategy"), "", "")
	cmd.PersistentFlags().StringVar(&req.VideoUrl, cfg.FlagNamer("VideoUrl"), "", "")
	cmd.PersistentFlags().StringVar(&req.UnitOfMeasures, cfg.FlagNamer("UnitOfMeasures"), "", "")
	flag.BytesBase64SliceVar(cmd.PersistentFlags(), &req.ImageUploads, cfg.FlagNamer("ImageUploads"), "")
	cmd.PersistentFlags().StringSliceVar(&req.ImageResourceIds, cfg.FlagNamer("ImageResourceIds"), nil, "")

	return cmd
}

func _SurveyServiceDeleteDiscoProjectCommand(cfg *client.Config) *cobra.Command {
	req := &IdRequest{}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("DeleteDiscoProject"),
		Short:  "DeleteDiscoProject RPC client",
		Long:   "",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "DeleteDiscoProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &IdRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.DeleteDiscoProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SurveyUserId, cfg.FlagNamer("SurveyUserId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountProjectId, cfg.FlagNamer("SysAccountProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SurveyProjectId, cfg.FlagNamer("SurveyProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountAccountId, cfg.FlagNamer("SysAccountAccountId"), "", "")
	cmd.PersistentFlags().StringVar(&req.DiscoProjectId, cfg.FlagNamer("DiscoProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountOrgId, cfg.FlagNamer("SysAccountOrgId"), "", "")

	return cmd
}

func _SurveyServiceGenTempIdCommand(cfg *client.Config) *cobra.Command {
	req := &emptypb.Empty{}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("GenTempId"),
		Short:  "GenTempId RPC client",
		Long:   "TempIdRequest",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "GenTempId"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &emptypb.Empty{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GenTempId(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	return cmd
}
