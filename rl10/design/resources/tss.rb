module Resources
  class TSS
    include Praxis::ResourceDefinition

    description 'Manipulate the TSS proxy'
    media_type 'text/plain'

    routing do
      prefix '/rll/tss'
    end

    action :put_control do
      description 'Control the TSS monitoring'
      routing { put '/control' }
      params do
        attribute :tss_id, String, required: false
        attribute :enable_monitoring, Attributor::Boolean, required: false
      end
      response :ok
    end

    action :get_hostname do
      description 'Get the TSS hostname to proxy'
      routing { get '/hostname' }
      response :ok
    end

    action :put_hostname do
      description 'Set the TSS hostname to proxy'
      routing { put '/hostname' }
      params do
        attribute :hostname, String, required: true
      end
      response :ok
    end
  end
end
