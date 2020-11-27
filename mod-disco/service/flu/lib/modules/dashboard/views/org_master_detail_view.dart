import 'dart:typed_data';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:mod_disco/core/core.dart';
import 'package:mod_disco/modules/dashboard/view_model/dashboard_view_model.dart';

// import 'package:mod_disco/modules/org_manager/orgs/data/org_model.dart';
import 'package:provider_architecture/provider_architecture.dart';
import 'package:responsive_builder/responsive_builder.dart';
import 'package:sys_core/sys_core.dart';
import 'package:sys_share_sys_account_service/rpc/v2/sys_account_models.pb.dart';

class OrgMasterDetailView extends StatelessWidget {
  final String id;
  final String orgId;

  OrgMasterDetailView({Key key, this.id = '', this.orgId = ''})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ViewModelProvider.withConsumer(
      viewModelBuilder: () => DashboardViewModel(),
      onModelReady: (DashboardViewModel model) async {
        if (model.orgs == null || model.orgs.isEmpty) {
          await model.fetchInitialProjects();
        }
      },
      builder: (context, DashboardViewModel model, child) =>
          NewGetCourageMasterDetail<Org, Project>(
        enableSearchBar: true,
        parentId: orgId,
        childId: id,
        items: model.orgs,
        labelBuilder: (item) => item.name,
        imageBuilder: (item) => item.logo,
        routeWithIdPlaceholder: Modular.get<Paths>().dashboardId,
        detailsBuilder: (context, parentId, childId, isFullScreen) {
          model.getSelectedProjectAndDetails(parentId, childId);
          return _getDetailsView(
              context, parentId, childId, isFullScreen, model);
        },
        noItemsAvailable: Center(
          child: Text(
            ModDiscoLocalizations.of(context).translate('noCampaigns'),
          ),
        ),
        noItemsSelected: Center(
            child: Text(ModDiscoLocalizations.of(context)
                .translate('noItemsSelected'))),
        disableBackButtonOnNoItemSelected: false,
        masterAppBarTitle:
            Text(ModDiscoLocalizations.of(context).translate('selectCampaign')),
        isLoadingMoreItems: model.isLoading,
        fetchNextItems: model.fetchNextProjects,
        hasMoreItems: model.hasMoreItems,
        searchFunction: model.searchProjects,
        resetSearchFunction: model.onResetSearchProjects,
        itemChildren: (org) => org.projects,
        childBuilder: (project, id) => <Widget>[
          ...[
            SizedBox(width: 16),
            CircleAvatar(
              radius: 20,
              backgroundImage: MemoryImage(
                Uint8List.fromList(project.logo),
              ),
            ),
          ],
          SizedBox(width: 16),
          //logic taken from ListTile
          Expanded(
            child: Text(
              project.name,
              textAlign: TextAlign.center,
              style: Theme.of(context).textTheme.subtitle1.merge(
                    TextStyle(
                      color: project.id != id
                          ? Theme.of(context).textTheme.subtitle1.color
                          : Theme.of(context).accentColor,
                    ),
                  ),
            ),
          ),
          SizedBox(width: 30),
        ],
      ),
    );
  }

  Widget _getDetailsView(BuildContext context, String parentId, String childId,
      bool isFullScreen, DashboardViewModel model) {
    return ViewModelProvider.withConsumer(
      viewModelBuilder: () => model,
      builder: (context, DashboardViewModel model, child) => ResponsiveBuilder(
        builder: (context, sizingInfo) {
          return Scaffold(
            appBar: AppBar(
              // iconTheme: Theme.of(context).iconTheme,
              automaticallyImplyLeading: isFullScreen,
              title:
                  Text(model.orgs.firstWhere((org) => org.id == parentId).name),
              // this the mock data
              actions: <Widget>[
                IconButton(
                    tooltip: "Copy Link",
                    icon: Icon(Icons.link),
                    onPressed: () async {
                      String link =
                          "${Modular.get<EnvConfig>().url}/${Modular.get<Paths>().dashboardId.replaceFirst("/", "").replaceAll(":id", "${parentId}")}";
                      // ${Modular.get<Paths>().org.replaceAll(":id", "$index")
                      print(Modular.get<Paths>().baseRoute);
                      await Clipboard.setData(new ClipboardData(text: link));
                      print(link);
                      print(Modular.get<Paths>()
                          .dashboardId
                          .replaceFirst("/", "")
                          .replaceAll(":id", "1"));
                    })
              ],
            ),
            drawer: (sizingInfo.screenSize.width > 1100)
                ? null
                // : Drawer(child: FilterPane(sizingInfo: sizingInfo)),
                : Drawer(child: Container(child: Text('Filter Pane here'))),
            body: Container(),
            // body: Column(
            //   crossAxisAlignment: CrossAxisAlignment.start,
            //   children: <Widget>[
            //     (sizingInfo.screenSize.width <= 768)
            //         ? IconButton(
            //             icon: Icon(Icons.arrow_back),
            //             onPressed: () {
            //               return Modular.to.pop(false);
            //             },
            //           )
            //         : Offstage(),
            //     Expanded(
            //       child: Row(
            //         crossAxisAlignment: CrossAxisAlignment.start,
            //         children: <Widget>[
            //           (sizingInfo.screenSize.width > 1100)
            //               // ? FilterPane(sizingInfo: sizingInfo)
            //               ? Container(child: Text('FilterPane here'))
            //               : Offstage(),
            //           SizedBox(width: 16),
            //           // Expanded(
            //           //     child: DataPane(
            //           //   sizingInfo: sizingInfo,
            //           //   model: model,
            //           // )),
            //         ],
            //       ),
            //     ),
            //   ],
            // ),
          );
        },
      ),
    );
  }
}
