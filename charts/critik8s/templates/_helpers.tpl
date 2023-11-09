{{/*
Expand the name of the chart.
*/}}
{{- define "amqp-proxy.name" -}}
{{- default .Chart.Name .Values.amqpProxy.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "amqp-proxy.fullname" -}}
{{- if .Values.amqpProxy.fullnameOverride }}
{{- .Values.amqpProxy.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.amqpProxy.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "amqp-proxy.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "amqp-proxy.labels" -}}
helm.sh/chart: {{ include "amqp-proxy.chart" . }}
{{ include "amqp-proxy.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "amqp-proxy.selectorLabels" -}}
app.kubernetes.io/name: {{ include "amqp-proxy.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "amqp-proxy.serviceAccountName" -}}
{{- if .Values.amqpProxy.serviceAccount.create }}
{{- default (include "amqp-proxy.fullname" .) .Values.amqpProxy.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.amqpProxy.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Expand the name of the chart.
*/}}
{{- define "data-collector.name" -}}
{{- default .Chart.Name .Values.dataCollector.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "data-collector.fullname" -}}
{{- if .Values.dataCollector.fullnameOverride }}
{{- .Values.dataCollector.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.dataCollector.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "data-collector.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "data-collector.labels" -}}
helm.sh/chart: {{ include "data-collector.chart" . }}
{{ include "data-collector.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "data-collector.selectorLabels" -}}
app.kubernetes.io/name: {{ include "data-collector.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "data-collector.serviceAccountName" -}}
{{- if .Values.dataCollector.serviceAccount.create }}
{{- default (include "data-collector.fullname" .) .Values.dataCollector.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.dataCollector.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Expand the name of the chart.
*/}}
{{- define "monitor-backend.name" -}}
{{- default .Chart.Name .Values.monitorBackend.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "monitor-backend.fullname" -}}
{{- if .Values.monitorBackend.fullnameOverride }}
{{- .Values.monitorBackend.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.monitorBackend.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "monitor-backend.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "monitor-backend.labels" -}}
helm.sh/chart: {{ include "monitor-backend.chart" . }}
{{ include "monitor-backend.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "monitor-backend.selectorLabels" -}}
app.kubernetes.io/name: {{ include "monitor-backend.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "monitor-backend.serviceAccountName" -}}
{{- if .Values.monitorBackend.serviceAccount.create }}
{{- default (include "monitor-backend.fullname" .) .Values.monitorBackend.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.monitorBackend.serviceAccount.name }}
{{- end }}
{{- end }}
