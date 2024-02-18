{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "project.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Expand the name of project .
*/}}
{{- define "project.name" -}}
{{- .Values.global.name | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
rocktemplateAgent Common labels
*/}}
{{- define "project.rocktemplateAgent.labels" -}}
helm.sh/chart: {{ include "project.chart" . }}
{{ include "project.rocktemplateAgent.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
rocktemplateAgent Common labels
*/}}
{{- define "project.rocktemplateController.labels" -}}
helm.sh/chart: {{ include "project.chart" . }}
{{ include "project.rocktemplateController.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
rocktemplateAgent Selector labels
*/}}
{{- define "project.rocktemplateAgent.selectorLabels" -}}
app.kubernetes.io/name: {{ include "project.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/component: {{ .Values.rocktemplateAgent.name | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
rocktemplateAgent Selector labels
*/}}
{{- define "project.rocktemplateController.selectorLabels" -}}
app.kubernetes.io/name: {{ include "project.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/component: {{ .Values.rocktemplateController.name | trunc 63 | trimSuffix "-" }}
{{- end }}


{{/* vim: set filetype=mustache: */}}
{{/*
Renders a value that contains template.
Usage:
{{ include "tplvalues.render" ( dict "value" .Values.path.to.the.Value "context" $) }}
*/}}
{{- define "tplvalues.render" -}}
    {{- if typeIs "string" .value }}
        {{- tpl .value .context }}
    {{- else }}
        {{- tpl (.value | toYaml) .context }}
    {{- end }}
{{- end -}}




{{/*
Return the appropriate apiVersion for poddisruptionbudget.
*/}}
{{- define "capabilities.policy.apiVersion" -}}
{{- if semverCompare "<1.21-0" .Capabilities.KubeVersion.Version -}}
{{- print "policy/v1beta1" -}}
{{- else -}}
{{- print "policy/v1" -}}
{{- end -}}
{{- end -}}

{{/*
Return the appropriate apiVersion for deployment.
*/}}
{{- define "capabilities.deployment.apiVersion" -}}
{{- if semverCompare "<1.14-0" .Capabilities.KubeVersion.Version -}}
{{- print "extensions/v1beta1" -}}
{{- else -}}
{{- print "apps/v1" -}}
{{- end -}}
{{- end -}}


{{/*
Return the appropriate apiVersion for RBAC resources.
*/}}
{{- define "capabilities.rbac.apiVersion" -}}
{{- if semverCompare "<1.17-0" .Capabilities.KubeVersion.Version -}}
{{- print "rbac.authorization.k8s.io/v1beta1" -}}
{{- else -}}
{{- print "rbac.authorization.k8s.io/v1" -}}
{{- end -}}
{{- end -}}

{{/*
return the rocktemplateAgent image
*/}}
{{- define "project.rocktemplateAgent.image" -}}
{{- $registryName := .Values.rocktemplateAgent.image.registry -}}
{{- $repositoryName := .Values.rocktemplateAgent.image.repository -}}
{{- if .Values.global.imageRegistryOverride }}
    {{- printf "%s/%s" .Values.global.imageRegistryOverride $repositoryName -}}
{{ else if $registryName }}
    {{- printf "%s/%s" $registryName $repositoryName -}}
{{- else -}}
    {{- printf "%s" $repositoryName -}}
{{- end -}}
{{- if .Values.rocktemplateAgent.image.digest }}
    {{- print "@" .Values.rocktemplateAgent.image.digest -}}
{{- else if .Values.global.imageTagOverride -}}
    {{- printf ":%s" .Values.global.imageTagOverride -}}
{{- else if .Values.rocktemplateAgent.image.tag -}}
    {{- printf ":%s" .Values.rocktemplateAgent.image.tag -}}
{{- else -}}
    {{- printf ":v%s" .Chart.AppVersion -}}
{{- end -}}
{{- end -}}


{{/*
return the rocktemplateController image
*/}}
{{- define "project.rocktemplateController.image" -}}
{{- $registryName := .Values.rocktemplateController.image.registry -}}
{{- $repositoryName := .Values.rocktemplateController.image.repository -}}
{{- if .Values.global.imageRegistryOverride }}
    {{- printf "%s/%s" .Values.global.imageRegistryOverride $repositoryName -}}
{{ else if $registryName }}
    {{- printf "%s/%s" $registryName $repositoryName -}}
{{- else -}}
    {{- printf "%s" $repositoryName -}}
{{- end -}}
{{- if .Values.rocktemplateController.image.digest }}
    {{- print "@" .Values.rocktemplateController.image.digest -}}
{{- else if .Values.global.imageTagOverride -}}
    {{- printf ":%s" .Values.global.imageTagOverride -}}
{{- else if .Values.rocktemplateController.image.tag -}}
    {{- printf ":%s" .Values.rocktemplateController.image.tag -}}
{{- else -}}
    {{- printf ":v%s" .Chart.AppVersion -}}
{{- end -}}
{{- end -}}


{{/*
generate the CA cert
*/}}
{{- define "generate-ca-certs" }}
    {{- $ca := genCA "spidernet.io" (.Values.rocktemplateController.tls.auto.caExpiration | int) -}}
    {{- $_ := set . "ca" $ca -}}
{{- end }}
