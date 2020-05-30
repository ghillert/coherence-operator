<doc-view>

<h2 id="_publish_metrics">Publish Metrics</h2>
<div class="section">
<p>Since version 12.2.1.4 Coherence has had the ability to expose a http endpoint that can be used to scrape metrics.
This would typically be used to expose metrics to something like Prometheus.</p>

<div class="admonition note">
<p class="admonition-inline">The metrics endpoint is <strong>disabled</strong> by default in Coherence clusters but can be enabled and configured by
setting the relevant fields in the <code>Coherence</code> CRD.</p>
</div>
<p>The example below shows how to enable and access Coherence metrics.</p>

<p>Once the metrics port has been exposed, for example via a load balancer or port-forward command, the metrics
endpoint is available at <code><a id="" title="" target="_blank" href="http://host:port/metrics">http://host:port/metrics</a></code>.</p>

<p>See the <a id="" title="" target="_blank" href="https://docs.oracle.com/en/middleware/standalone/coherence/14.1.1.0/manage/using-coherence-metrics.html">Using Coherence Metrics</a>
documentation for full details on the available metrics.</p>

<div class="admonition note">
<p class="admonition-inline">Note: Use of metrics is available only when using the operator with clusters running Coherence 12.2.1.4 or later version.</p>
</div>

<h3 id="_deploy_coherence_with_metrics_enabled">Deploy Coherence with Metrics Enabled</h3>
<div class="section">
<p>To deploy a <code>Coherence</code> resource with metrics enabled and exposed on a port, the simplest yaml
would look like this:</p>

<markup
lang="yaml"
title="metrics-cluster.yaml"
>apiVersion: coherence.oracle.com/v1
kind: Coherence
metadata:
  name: metrics-cluster
spec:
  coherence:
    metrics:
      enabled: true     <span class="conum" data-value="1" />
  ports:
    - name: metrics  <span class="conum" data-value="2" /></markup>

<ul class="colist">
<li data-value="1">Setting the <code>coherence.metrics.enabled</code> field to <code>true</code> will enable metrics</li>
<li data-value="2">To expose metrics via a <code>Service</code> it is added to the <code>ports</code> list.
The <code>metrics</code> port is a special case where the <code>port</code> number is optional so in this case metrics
will bind to the default port <code>9612</code>.
(see <router-link to="/ports/020_container_ports">Exposing Ports</router-link> for details)</li>
</ul>
<p>To expose metrics on a different port the alternative port value can be set in the <code>coherence.metrics</code>
section, for example:</p>

<markup
lang="yaml"
title="metrics-cluster.yaml"
>apiVersion: coherence.oracle.com/v1
kind: Coherence
metadata:
  name: metrics-cluster
spec:
  coherence:
    metrics:
      enabled: true
      port: 8080      <span class="conum" data-value="1" />
  ports:
    - name: metrics</markup>

<ul class="colist">
<li data-value="1">metrics will now be exposed on port <code>8080</code></li>
</ul>
</div>

<h3 id="_port_forward_the_metrics_port">Port-forward the Metrics Port</h3>
<div class="section">
<p>After installing the basic <code>metrics-cluster.yaml</code> from the first example above there would be a three member
Coherence cluster installed into Kubernetes.</p>

<p>For example, the cluster can be installed with <code>kubectl</code></p>

<markup
lang="bash"

>kubectl -n coherence-test create -f metrics-cluster.yaml

coherence.coherence.oracle.com/metrics-cluster created</markup>

<p>The <code>kubectl</code> CLI can be used to list <code>Pods</code> for the cluster:</p>

<markup
lang="bash"

>kubectl -n coherence-test get pod -l coherenceCluster=metrics-cluster

NAME                   READY   STATUS    RESTARTS   AGE
metrics-cluster-0   1/1     Running   0          36s
metrics-cluster-1   1/1     Running   0          36s
metrics-cluster-2   1/1     Running   0          36s</markup>

<p>In a test or development environment the simplest way to reach an exposed port is to use the <code>kubectl port-forward</code> command.
For example to connect to the first <code>Pod</code> in the deployment:</p>

<markup
lang="bash"

>kubectl -n coherence-test port-forward metrics-cluster-0 9612:9612

Forwarding from [::1]:9612 -&gt; 9612
Forwarding from 127.0.0.1:9612 -&gt; 9612</markup>

</div>

<h3 id="_access_the_metrics_endpoint">Access the Metrics Endpoint</h3>
<div class="section">
<p>Now that a port has been forwarded from localhost to a <code>Pod</code> in the cluster the metrics endpoint can be accessed.</p>

<p>Issue the following <code>curl</code> command to access the REST endpoint:</p>

<markup
lang="bash"

>curl http://127.0.0.1:9612/metrics</markup>

</div>
</div>
</doc-view>
